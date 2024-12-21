import ConflictError from "@common/exceptions/ConflictError";
import MongoServerError from "@common/exceptions/MongoServerError";
import UnauthorizedError from "@common/exceptions/UnauthorizeError";
import { JWTContext } from "@common/types/extends/JWTContext";
import { User } from "@domain/User";
import * as userRepository from "@repository/user";

export const signIn = async (context: JWTContext, payload: User): Promise<string> => {
  const user = await userRepository.findByEmail(payload.email);
  if (!user) {
    throw new UnauthorizedError("User not found");
  }
  const isPasswordValid = await user.comparePassword(payload.password);
  if (!isPasswordValid) {
    throw new UnauthorizedError("Invalid password");
  }

  return await context.jwt.sign({ id: user.id, email: user.email });
};

export const signUp = async (context: JWTContext, payload: User): Promise<string> => {
  try {
    const user = await userRepository.findByEmail(payload.email);
    if (user) {
      throw new ConflictError("User already exists");
    }
    const createdUser = await userRepository.create(payload);

    return await context.jwt.sign({ id: createdUser.id, email: createdUser.email });
  } catch (e) {
    const error = e as MongoServerError;
    if (error.name === "MongoServerError" && error.code === 11000) {
      throw new ConflictError("User already exists");
    }
    throw error;
  }
};
