import { Elysia } from "elysia";

export default (app: Elysia) =>
  // @ts-expect-error
  app.derive(async ({ jwt, cookie }) => {
    const accessToken = cookie.session.value;

    if (!accessToken) {
      return;
    }

    const currentUser = await jwt.verify(accessToken);

    return {
      currentUser,
    };
  });
