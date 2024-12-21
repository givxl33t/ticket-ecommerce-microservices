import { User } from "@infrastructure/database/schema/User";

export async function findByEmail(email: string): Promise<User | null> {
  return await User.findOne({ email });
}

export async function findById(id: string): Promise<User | null> {
  return await User.findById(id);
}

export async function findAll(): Promise<User[]> {
  return await User.find();
}

export async function create(payload: unknown): Promise<User> {
  return await User.create(payload);
}
