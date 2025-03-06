import { t } from "elysia";

export const authDto = {
    body: t.Object({
        email: t.String({format: 'email'}),
        password: t.String(),
    })
}