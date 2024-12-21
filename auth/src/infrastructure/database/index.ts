import * as mongoose from "mongoose";

import config from "@/config";

const { uri } = config.db;

export const connect = async () => {
  try {
    const res = await mongoose.connect(uri, {
      autoIndex: true,
    });

    console.log("🚀 MongoDB successfully connected: ", res.connection.name);
  } catch (err) {
    console.log("❌ MongoDB connection error: ", err);
  }
};
