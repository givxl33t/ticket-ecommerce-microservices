import { Publisher, Subjects, TicketUpdatedEvent } from "@romen-tix-micro/common";

export class TicketUpdatedPublisher extends Publisher<TicketUpdatedEvent> {
  subject: Subjects.TicketUpdated = Subjects.TicketUpdated;
}