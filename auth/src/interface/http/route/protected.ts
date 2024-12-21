import * as authController from "@interface/http/controller/auth";
import { Elysia } from "elysia";

import requireAuth from "@/src/infrastructure/plugins/requireAuth";

export default (app: Elysia) =>
  app.use(requireAuth).post("/signout", authController.signOut, {
    detail: {
      tags: ["Auth"],
      security: [{ cookieAuth: [] }],
    },
  });
