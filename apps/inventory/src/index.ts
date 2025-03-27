import { Elysia } from "elysia";
import { swagger } from '@elysiajs/swagger'
import { cors } from "@elysiajs/cors";

import { authController } from "./auth/auth.controller";
import { inventoryController } from "./inventory/inventory.controller";

const app = new Elysia()
.use(swagger())
.use(cors())
.use(authController)
.use(inventoryController)
.get("/health", () => {
  return {
    status: "ok",
  };
})
.listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.url}`
);
