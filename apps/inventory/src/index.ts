import { Elysia } from "elysia";
import { swagger } from '@elysiajs/swagger'

import { authController } from "./auth/auth.controller";
import { inventoryController } from "./inventory/inventory.controller";

const app = new Elysia()
.use(swagger())
.use(authController)
.use(inventoryController)
.listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.url}`
);
