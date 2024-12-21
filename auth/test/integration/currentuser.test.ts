/* eslint-disable prettier/prettier */
/* eslint-disable sonarjs/no-hardcoded-passwords */
import app from "@infrastructure/app";
import { User } from "@infrastructure/database/schema/User";
import { afterAll, beforeEach, describe, expect, it } from "bun:test";

import { truncate } from "./utils/databaseHelper";
import { getRequest, postRequest } from "./utils/httpRequest";

describe("Current User Route", () => {
  afterAll(async () => {
    await truncate({ User });
  });

  beforeEach(async () => {
    await truncate({ User });
  });

  it("responds with details about the current user", async () => {
    const signUpResponse = await app.handle(
      postRequest("/api/users/signup", {
        email: "test@test.com",
        password: "password",
      }),
    );

    const responseCookie = signUpResponse.headers.getSetCookie();

    const response = await app
      .handle(getRequest("/api/users/currentuser", { headers: { Cookie: responseCookie } }))
      .then((res: Response) => res.json());

    expect(response.data.email).toEqual("test@test.com");
  });

  it("response with null if not authenticated", async () => {
    const response = await app.handle(getRequest("/api/users/currentuser")).then((res: Response) => res.json());
    
    expect(response.data).toBeNull();
  });
});
