import Elysia from "elysia";
import process from "process";
import * as colors from "yoctocolors";

// returns the duration message
export const durationString = (beforeTime: bigint): string => {
  const now = process.hrtime.bigint();
  const timeDifference = now - beforeTime;
  const nanoseconds = Number(timeDifference);

  let timeMessage: string = "";

  if (nanoseconds >= 1e9) {
    const seconds = (nanoseconds / 1e9).toFixed(2);
    timeMessage = `| ${seconds}s`;
  } else if (nanoseconds >= 1e6) {
    const durationInMilliseconds = (nanoseconds / 1e6).toFixed(0);

    timeMessage = `| ${durationInMilliseconds}ms`;
  } else if (nanoseconds >= 1e3) {
    const durationInMilliseconds = (nanoseconds / 1e3).toFixed(0);

    timeMessage = `| ${durationInMilliseconds}μs`;
  } else {
    timeMessage = `| ${nanoseconds}ns`;
  }
  return timeMessage;
};

// returns the method message
export const methodString = (method: string): string => {
  switch (method) {
    case "GET":
      return colors.white("GET");
    case "POST":
      return colors.yellow("POST");
    case "PUT":
      return colors.blue("PUT");
    case "DELETE":
      return colors.red("DELETE");
    case "PATCH":
      return colors.green("PATCH");
    case "OPTIONS":
      return colors.gray("OPTIONS");
    case "HEAD":
      return colors.magenta("HEAD");
    default:
      return method;
  }
};

export default (app: Elysia) =>
  app
    .state({ beforeTime: process.hrtime.bigint(), as: "global" })
    .onRequest((ctx) => {
      ctx.store.beforeTime = process.hrtime.bigint();
    })
    .onBeforeHandle({ as: "global" }, (ctx) => {
      ctx.store.beforeTime = process.hrtime.bigint();
    })
    .onAfterHandle({ as: "global" }, ({ request, store }) => {
      const logStr: string[] = [];

      logStr.push(methodString(request.method));
      logStr.push(new URL(request.url).pathname);

      const beforeTime: bigint = store.beforeTime;

      logStr.push(durationString(beforeTime));

      console.log(logStr.join(" "));
    })
    .onError({ as: "global" }, ({ request, error, store }) => {
      const logStr: string[] = [];

      logStr.push(colors.red(methodString(request.method)));
      logStr.push(new URL(request.url).pathname);
      logStr.push(colors.red("Error"));

      if ("status" in error) {
        logStr.push(String(error.status));
      }

      logStr.push(error.message);

      const beforeTime: bigint = store.beforeTime;
      logStr.push(durationString(beforeTime));

      console.log(logStr.join(" "));
    });
