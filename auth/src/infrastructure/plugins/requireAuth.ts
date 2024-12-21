import UnauthorizedError from "@common/exceptions/UnauthorizeError";
import { Elysia } from "elysia";

export default (app: Elysia) =>
  // @ts-expect-error
  app.derive(async ({ currentUser }) => {
    if (!currentUser) {
      throw new UnauthorizedError("Unauthorized");
    }

    return {
      currentUser,
    };
  });
