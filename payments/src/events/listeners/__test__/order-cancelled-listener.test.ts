import { OrderCancelledListener } from "../order-cancelled-listener";
import { natsWrapper } from "../../../nats-wrapper";
import { OrderCancelledEvent, OrderStatus } from "@romen-tix-micro/common";
import mongoose from 'mongoose';
import { Message } from 'node-nats-streaming';
import { Order } from "../../../models/order";

const setup = async () => {
  const listener = new OrderCancelledListener(natsWrapper.client);

  const orderId = new mongoose.Types.ObjectId().toHexString();
  const order = Order.build({
    id: orderId,
    version: 0,
    status: OrderStatus.Created,
    userId: 'asdf',
    price: 10,
  });
  await order.save();

  const data: OrderCancelledEvent['data'] = {
    id: orderId,
    version: 1,
    ticket: {
      id: 'asdf',
    }
  };

  // @ts-ignore
  const msg: Message = {
    ack: jest.fn(),
  };

  return { listener, data, msg, order, orderId };
};

it('updates the status of the order', async () => {
  const { listener, data, msg, orderId } = await setup();

  await listener.onMessage(data, msg);

  const updatedOrder = await Order.findById(orderId);

  expect(updatedOrder!.status).toEqual(OrderStatus.Cancelled);
});

it('acks the message', async () => {
  const { listener, data, msg } = await setup();

  await listener.onMessage(data, msg);

  expect(msg.ack).toHaveBeenCalled();
});