import { MongoMemoryServer } from 'mongodb-memory-server';
import mongoose from 'mongoose';
import jwt from 'jsonwebtoken'

let mongo: any;

declare global {
  var signup: (id?: string) =>string[];
}

jest.mock('../nats-wrapper');

process.env.STRIPE_KEY = 'sk_test_51NuxwkEPFj7I6FUPMJqh9CQ6Woc8Rm5bSvYpMT6zCBISCVz3XLLQg9BFtUvPh5qqm6kosy5FaWo89JqFBaAsz31T00KI1CX9ke'

beforeAll(async () => {
  process.env.JWT_KEY = 'asdfasdf';

  // Set up a fake MongoDB server in memory
  mongo = await MongoMemoryServer.create();
  const mongoUri = mongo.getUri();

  // Connect mongoose to the fake MongoDB server
  await mongoose.connect(mongoUri, {});
});

beforeEach(async () => {
  jest.clearAllMocks();
  // Reset all data between tests
  const collections = await mongoose.connection.db.collections();

  for (let collection of collections) {
    await collection.deleteMany({});
  }
});

afterAll(async () => {
  if (mongo) {
    await mongo.stop();
  }
  await mongoose.connection.close();
});

global.signup = (id?: string) => {
  const payload = {
    id: id || new mongoose.Types.ObjectId().toHexString(),
    email: 'test@test.com',
  };

  const token = jwt.sign(payload, process.env.JWT_KEY!);

  const sessionJSON = JSON.stringify({ jwt: token });

  const base64 = Buffer.from(sessionJSON).toString('base64');

  return [`session=${base64}`];
};