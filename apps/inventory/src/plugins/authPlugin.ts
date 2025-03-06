import Elysia from "elysia";
import jwt from "@elysiajs/jwt";
import { prismaClient } from "../libs/prismaClient";

export const authPlugin = (app: Elysia) => {
    return app.use(
        jwt({
            name: "jwt",
            secret: Bun.env.JWT_SECRET as string,
        })
    ).derive(async ({jwt, cookie: {accessToken, refreshToken}, error}) => {
        if (!accessToken.value) {
            return error(401, {message: "Unauthorized"})
        }

        const jwtPayload = await jwt.verify(accessToken.value);
        if (!jwtPayload) {
            return error(403, {message: "Forbidden"})
        }

        const userId = jwtPayload.sub;
        const user = await prismaClient.user.findUnique({
            where: {
                id: userId
            }
        });
        if (!user) {
            return error(403, {message: "Forbidden"})
        }

        return { user }
    })
}
