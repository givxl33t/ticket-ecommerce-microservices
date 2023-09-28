import { Listener, OrderCancelledEvent, OrderStatus, Subjects } from "@romen-tix-micro/common";
import { Message } from "node-nats-streaming";
import { queueGroupName } from "./queue-group-name";
import { Order } from "../../models/order";

export class OrderCancelledListener extends Listener<OrderCancelledEvent> {
  subject: Subjects.OrderCancelled = Subjects.OrderCancelled;
  queueGroupName = queueGroupName;

  async onMessage(data: OrderCancelledEvent["data"], msg: Message) {
    // find the order
    const order = await Order.findByEvent(data);

    // if no order, throw error
    if (!order) {
      throw new Error('Order not found');
    }

    // mark the order as being cancelled
    order.set({ status: OrderStatus.Cancelled });

    // save the order
    await order.save();

    // ack the message
    msg.ack();
  }
}