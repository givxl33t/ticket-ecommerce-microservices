/* eslint-disable sonarjs/no-hardcoded-passwords */
import app from "@infrastructure/app";
import { User } from "@infrastructure/database/schema/User";
import { afterAll, beforeEach, describe, expect, it } from "bun:test";

import { truncate } from "./utils/databaseHelper";
import { postRequest } from "./utils/httpRequest";

describe("Sign Up Route", () => {
  afterAll(async () => {
    await truncate({ User });
  });

  beforeEach(async () => {
    await truncate({ User });
  });

  it("returns a 200 on successful signup", async () => {
    const response = await app.handle(
      postRequest("/api/users/signup", {
        email: "test@test.com",
        password: "password",
      }),
    );

    expect(response.status).toBe(200);
  });

  it("returns a 400 with an invalid email", async () => {
    const response = await app.handle(
      postRequest("/api/users/signup", {
        email: "testtest.com",
        password: "password",
      }),
    );

    expect(response.status).toBe(400);
  });

  it("returns a 400 with an invalid password", async () => {
    const response = await app.handle(
      postRequest("/api/users/signup", {
        email: "testtest.com",
        password: "p",
      }),
    );

    expect(response.status).toBe(400);
  });

  it("returns a 400 with missing email and password", async () => {
    const response1 = await app.handle(
      postRequest("/api/users/signup", {
        email: "testtest.com",
      }),
    );

    expect(response1.status).toBe(400);

    const response2 = await app.handle(postRequest("/api/users/signup", {}));

    expect(response2.status).toBe(400);
  });

  it("disallows duplicate emails", async () => {
    await app.handle(
      postRequest("/api/users/signup", {
        email: "test@test.com",
        password: "password",
      }),
    );

    const response = await app.handle(
      postRequest("/api/users/signup", {
        email: "test@test.com",
        password: "password",
      }),
    );

    expect(response.status).toBe(409);
  });

  it("sets a cookie after successful signup", async () => {
    const response = await app.handle(
      postRequest("/api/users/signup", {
        email: "test@test.com",
        password: "password",
      }),
    );

    expect(response.headers.getSetCookie()).toBeDefined();
  });
});
