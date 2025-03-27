import { initContract } from "@ts-rest/core";
import { z } from "zod";

const c = initContract()

export const CustomerContract = c.router({
    getCustomer: {
        method : "GET",
        path: "/",
        query: z.object({
    }
})