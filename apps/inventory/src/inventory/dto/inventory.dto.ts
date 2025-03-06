import { t } from "elysia";

export const inventoryCreateDto = {
  body: t.Object({
    name: t.String(),
    weight: t.Number(),
    remarks: t.String({ optional: true }),
  }),
};

export const inventoryUpdateDto = {
    params: t.Object({
        id: t.String(),
    }),
    body: t.Object({
        name: t.String(),
        weight: t.Number(),
        remarks: t.String({ optional: true }),
    })
}