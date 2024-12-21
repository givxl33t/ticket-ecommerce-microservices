import { Elysia, t } from "elysia";

export const authSchema = new Elysia().model({
  auth: t.Object({
    email: t.String({ format: "email", minLength: 1, maxLength: 255 }),
    password: t.String({ minLength: 1, maxLength: 255 }),
  }),
});
