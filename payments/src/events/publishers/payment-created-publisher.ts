import { Publisher, Subjects, PaymentCreatedEvent } from "@romen-tix-micro/common";

export class PaymentCreatedPublisher extends Publisher<PaymentCreatedEvent> {
  subject: Subjects.PaymentCreated = Subjects.PaymentCreated;
}