import { Elysia } from "elysia";
import { jwt } from "@elysiajs/jwt";

import { AuthService } from "./auth.service";
import { authPlugin } from "../plugins/authPlugin";
import { authDto } from "./dto/auth.dto";
import { prismaClient } from "../libs/prismaClient";

export const authController = new Elysia({ prefix: "/auth", detail: {
  tags: ["Auth"],
} })
  .decorate("prisma", prismaClient)
  .decorate("authService", new AuthService(prismaClient))
  .use(
    jwt({
      name: "jwt",
      secret: Bun.env.JWT_SECRET as string,
    })
  )

  .post(
    "/register",
    async ({ authService, body, error }) => {
      const { email } = body;

      const userExists = await authService.checkUserExists(email);
      if (userExists) {
        return error(400, { message: "User already exists" });
      }

      return authService.createUser(body);
    },
    authDto,
  )

  .post(
    "/login",
    async ({ body, jwt, cookie: { accessToken }, error }) => {
      const user = await prismaClient.user.findUnique({
        where: {
          email: body.email,
        },
      });

      if (!user) {
        return error(400, { message: "email or password is incorrect!" });
      }

      const isPasswordValid = await Bun.password.verify(
        body.password,
        user.password,
        'bcrypt'
      );
      if (!isPasswordValid) {
        return error(400, { message: "email or password is incorrect!" });
      }

      const accessTokenValue = await jwt.sign({
        sub: user.id,
        exp: Math.floor(Date.now() / 1000) + 30 * 24 * 60 * 60,
      });

      accessToken.set({
        value: accessTokenValue,
        httpOnly: true,
        maxAge: 7 * 24 * 60 * 60,
        path: "/",
      })

      return {
        message: "Login successful",
        data: {
          user: {
            id: user.id,
            email: user.email,
          },
          accessToken: accessTokenValue,
        }
      };
    },
    authDto
  )

  .use(authPlugin).get('/me', ({ user }) => user)
  .use(authPlugin).get('/logout', ({ cookie: { accessToken}}) => {
    accessToken.remove();
    return { message: "Logout successful" }
  });

