/* eslint-disable @typescript-eslint/no-explicit-any */
import path from "path";

import config from "@/config";

export const baseUrl = `http://${config.app.host}:${config.app.port}`;

export function getRequest(route: string, options: any = {}) {
  const fullPath = path.join(baseUrl, route);

  return new Request(fullPath, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      ...options.headers,
    },
  });
}

export function postRequest<Payload extends Record<string, unknown>>(
  route: string,
  payload: Payload,
  options: any = {},
) {
  const fullPath = path.join(baseUrl, route);

  return new Request(fullPath, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      ...options.headers,
    },
    body: JSON.stringify(payload),
  });
}
