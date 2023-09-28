import { Subjects, Publisher, OrderCancelledEvent } from "@romen-tix-micro/common";

export class OrderCancelledPublisher extends Publisher<OrderCancelledEvent> {
  subject: Subjects.OrderCancelled = Subjects.OrderCancelled;
}