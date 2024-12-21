/* eslint-disable sonarjs/no-hardcoded-passwords */
import app from "@infrastructure/app";
import { User } from "@infrastructure/database/schema/User";
import { afterAll, beforeEach, describe, expect, it } from "bun:test";

import { truncate } from "./utils/databaseHelper";
import { postRequest } from "./utils/httpRequest";

describe("Sign Out Route", () => {
  afterAll(async () => {
    await truncate({ User });
  });

  beforeEach(async () => {
    await truncate({ User });
  });

  it("clears the cookie after signing out", async () => {
    const signUpResponse = await app.handle(
      postRequest("/api/users/signup", {
        email: "test@test.com",
        password: "password",
      }),
    );

    const responseCookie = signUpResponse.headers.getSetCookie();

    const response = await app.handle(
      postRequest("/api/users/signout", {}, { headers: { Cookie: responseCookie } }),
    );

    expect(response.headers.getSetCookie()).toEqual([
      "session=; Max-Age=0; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT",
    ]);
  });
});
