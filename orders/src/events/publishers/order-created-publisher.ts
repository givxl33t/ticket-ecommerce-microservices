import { Publisher, OrderCreatedEvent, Subjects } from "@romen-tix-micro/common";

export class OrderCreatedPublisher extends Publisher<OrderCreatedEvent> {
  subject: Subjects.OrderCreated = Subjects.OrderCreated;
}