import { Context } from "elysia";

export interface JWTContext extends Context {
  jwt: {
    readonly sign: (payload: object) => Promise<string>;
    readonly verify: (jwt: string | undefined) => Promise<object>;
  };
}
