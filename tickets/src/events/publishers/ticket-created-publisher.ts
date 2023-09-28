import { Publisher, Subjects, TicketCreatedEvent } from "@romen-tix-micro/common";

export class TicketCreatedPublisher extends Publisher<TicketCreatedEvent> {
  subject: Subjects.TicketCreated = Subjects.TicketCreated;
}