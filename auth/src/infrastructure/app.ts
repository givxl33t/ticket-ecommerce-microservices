import * as database from "@infrastructure/database";
import documentationPlugin from "@infrastructure/plugins/documentation";
import errorPlugin from "@infrastructure/plugins/error";
import loggerPlugin from "@infrastructure/plugins/logger";
import securityPlugin from "@infrastructure/plugins/security";
import authRoutes from "@interface/http/route/auth";
import { Elysia } from "elysia";

import config from "@/config";
import protectedRoutes from "@/src/interface/http/route/protected";

const app = new Elysia();

database.connect();

if (config.app.env === "development") {
  app.use(loggerPlugin);
}

app
  .use(documentationPlugin)
  .use(securityPlugin)
  .use(errorPlugin)
  .group("/api/users", (group) => {
    return group.use(authRoutes).use(protectedRoutes);
  });

export default app;
