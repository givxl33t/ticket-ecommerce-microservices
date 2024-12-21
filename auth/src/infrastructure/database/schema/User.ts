import UnauthorizedError from "@common/exceptions/UnauthorizeError";
import mongoose from "mongoose";

const userSchema = new mongoose.Schema({
  email: {
    type: String,
    required: true,
    unique: true,
  },
  password: {
    type: String,
    required: true,
    select: false,
  },
});

userSchema.set("toJSON", {
  virtuals: true,
  transform: (doc, ret) => {
    ret.id = ret._id;
    delete ret._id;
    delete ret.password;
    delete ret.__v;
  },
});

userSchema.pre("save", function (next) {
  if (this.isModified("password")) {
    this.password = Bun.password.hashSync(this.password, {
      algorithm: "bcrypt",
    });
  }
  next();
});

userSchema.methods.comparePassword = async function (password: string) {
  const user = await User.findById(this._id).select("+password");
  if (!user) {
    throw new UnauthorizedError("User not found");
  }
  return Bun.password.verifySync(password, user.password);
};

type UserSchema = mongoose.InferSchemaType<typeof userSchema>;

export interface User extends UserSchema, mongoose.Document {
  comparePassword(password: string): Promise<boolean>;
}

export const User = mongoose.model<User>("User", userSchema);
