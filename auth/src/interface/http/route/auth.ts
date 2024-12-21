import currentUser from "@infrastructure/plugins/currentUser";
import { authSchema } from "@infrastructure/validation/authSchema";
import * as authController from "@interface/http/controller/auth";
import { Elysia } from "elysia";

export default (app: Elysia) =>
  app
    .use(authSchema)
    .use(currentUser)
    .get("/currentuser", authController.currentUser, {
      detail: {
        tags: ["Auth"],
      },
    })
    .post("/signup", authController.signUp, {
      body: "auth",
      detail: {
        tags: ["Auth"],
      },
    })
    .post("/signin", authController.signIn, {
      body: "auth",
      detail: {
        tags: ["Auth"],
      },
    });
