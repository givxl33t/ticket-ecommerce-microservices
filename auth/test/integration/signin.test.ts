/* eslint-disable sonarjs/no-hardcoded-passwords */
import app from "@infrastructure/app";
import { User } from "@infrastructure/database/schema/User";
import { afterAll, beforeEach, describe, expect, it } from "bun:test";

import { truncate } from "./utils/databaseHelper";
import { postRequest } from "./utils/httpRequest";

describe("Sign In Route", () => {
  afterAll(async () => {
    await truncate({ User });
  });

  beforeEach(async () => {
    await truncate({ User });
  });

  it("fails when an email that does not exist is supplied", async () => {
    const response = await app
      .handle(
        postRequest("/api/users/signin", {
          email: "adadada@mail.com",
          password: "password",
        }),
      )
      .then((res: Response) => res.json());

    expect(response).toMatchObject({
      code: 401,
      message: "User not found",
    });
  });

  it("fails when an incorrect password is supplied", async () => {
    const signUp = await app
      .handle(
        postRequest("/api/users/signup", {
          email: "test@test.com",
          password: "password123",
        }),
      )
      .then((res: Response) => res.json());

    expect(signUp).toMatchObject({
      message: "User created successfully",
    });

    const signIn = await app
      .handle(
        postRequest("/api/users/signin", {
          email: "test@test.com",
          password: "passwod",
        }),
      )
      .then((res: Response) => res.json());

    expect(signIn).toMatchObject({
      code: 401,
      message: "Invalid password",
    });
  });

  it("responds with a cookie when given a valid credentials", async () => {
    const signUp = await app
      .handle(
        postRequest("/api/users/signup", {
          email: "cat@cat.com",
          password: "password123",
        }),
      )
      .then((res: Response) => res.json());

    expect(signUp).toMatchObject({
      message: "User created successfully",
    });

    const signIn = await app.handle(
      postRequest("/api/users/signin", {
        email: "cat@cat.com",
        password: "password123",
      }),
    );

    const responseCookies = signIn.headers.getSetCookie();

    expect(responseCookies).toBeDefined();
  });
});
