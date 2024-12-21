import ConflictError from "@common/exceptions/ConflictError";
import UnauthorizedError from "@common/exceptions/UnauthorizeError";
import ErrorResponse from "@common/types/generics/ErrorResponse";
import { Elysia } from "elysia";
import { StatusCodes } from "http-status-codes";

import config from "@/config";

export default (app: Elysia) =>
  app.error({ ConflictError, UnauthorizedError }).onError((handler): ErrorResponse<number> => {
    if (config.app.env === "development") {
      console.error(handler.error?.stack);
    }

    if (handler.error instanceof ConflictError || handler.error instanceof UnauthorizedError) {
      handler.set.status = handler.error.status;

      return {
        message: handler.error.message,
        code: handler.error.status,
      };
    }

    if (handler.code === "NOT_FOUND") {
      handler.set.status = StatusCodes.NOT_FOUND;
      return {
        message: "Not Found!",
        code: handler.set.status,
      };
    }

    // overwrite 422 default code for validation
    if (handler.code === "VALIDATION") {
      handler.set.status = StatusCodes.BAD_REQUEST;
      return {
        message: JSON.parse(handler.error.message),
        code: handler.set.status,
      };
    }

    handler.set.status = StatusCodes.SERVICE_UNAVAILABLE;

    return {
      message: "Server Error!",
      code: handler.set.status,
    };
  });
