import TicketShow from "./TicketShow";
import axios from "axios";
import { headers } from "next/headers";

export interface ITicket {
  title: string;
  price: number;
  userId: string;
  version: number;
  id: string;
}

export default async function Page({ params }: { params: { ticketId: string } }) {
  const { ticketId } = params;
  const ticket: ITicket = await axios
    .get(`${process.env.API_GATEWAY_URL}/api/tickets/${ticketId}`, {
      headers: {
        ...Object.fromEntries(headers().entries()),
      },
    })
    .then((res) => res.data);

  return <TicketShow ticket={{ ...ticket }} />;
}
