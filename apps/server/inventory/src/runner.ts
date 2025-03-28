import type { SpawnOptions } from "bun";
import { execSync } from "child_process";

const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

const spawnOptions: SpawnOptions.OptionsObject = {
    stdin: "inherit",
    stdout: "inherit",
    stderr: "inherit",
}

const run = async () => {
    try {
        const bunPath = execSync("which bun").toString().trim();
        
        console.log("Starting main application...");
        Bun.spawn([bunPath, "run", "src/index.ts"], spawnOptions);
        
        console.log("Waiting for RabbitMQ to be ready...");
        await delay(5000);
        
        console.log("Starting RabbitMQ consumer...");
        Bun.spawn([bunPath, "run", "src/rabbitmq/consumer.ts"], spawnOptions);

        process.on("SIGINT", () => {
            console.log("SIGINT received, shutting down...");
            process.exit(0);
        });

        process.on("SIGTERM", () => {
            console.log("SIGTERM received, shutting down...");
            process.exit(0);
        });
    } catch (error) {
        console.error("Error running services:", error);
        process.exit(1);
    }
}

run();