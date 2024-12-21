import { JWTContext } from "@common/types/extends/JWTContext";
import { UserContext } from "@common/types/extends/UserContext";
import SuccessResponse from "@common/types/generics/SuccessResponse";
import { User } from "@domain/User";
import { JWTPayload } from "@model/Auth";
import * as authUsecase from "@usecase/auth";
import { Context } from "elysia";

export const currentUser = async (
  context: UserContext,
): Promise<SuccessResponse<JWTPayload | null>> => {
  return {
    message: "User details fetched successfully!",
    data: context.currentUser ? context.currentUser : null,
  };
};

export const signUp = async (context: JWTContext): Promise<SuccessResponse<string>> => {
  const payload = context.body as User;
  const token = await authUsecase.signUp(context, payload);

  context.cookie.session.set({
    value: token,
    httpOnly: false,
    secure: false,
    expires: new Date(Date.now() + 1000 * 60 * 60 * 24 * 7),
    maxAge: 1000 * 60 * 60 * 24 * 7,
  });

  return {
    message: "User created successfully",
  };
};

export const signIn = async (context: JWTContext): Promise<SuccessResponse<string>> => {
  const payload = context.body as User;
  const token = await authUsecase.signIn(context, payload);

  context.cookie.session.set({
    value: token,
    httpOnly: false,
    secure: false,
    expires: new Date(Date.now() + 1000 * 60 * 60 * 24 * 7),
    maxAge: 1000 * 60 * 60 * 24 * 7,
  });

  return {
    message: "User logged in successfully",
  };
};

export const signOut = async (context: Context): Promise<SuccessResponse<string>> => {
  context.cookie.session.remove();

  return {
    message: "User logged out successfully",
  };
};
