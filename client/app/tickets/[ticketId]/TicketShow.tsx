/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useRouter } from "next/navigation";
import { Button, Typography, Box } from "@mui/material";
import useRequest from "@/hooks/useRequest";
import { ITicket } from "./page";

export default function TicketShow({ ticket }: { ticket: ITicket }) {
  const router = useRouter();
  const { doRequest, errors } = useRequest({
    url: "/api/orders",
    method: "post",
    body: { ticketId: ticket.id },
    onSuccess: (order: any) => {
      router.push(`/orders/${order.id}`);
    },
  });

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        {ticket.title || "Ticket Details"}
      </Typography>
      <Typography variant="h6">Price: {ticket.price || "$0.00"}</Typography>
      {errors && <Box sx={{ width: "100%", mt: 2 }}>{errors}</Box>}
      <Button
        variant="contained"
        color="primary"
        onClick={() => doRequest()}
        style={{ marginTop: "1rem" }}
      >
        Purchase
      </Button>
    </Box>
  );
}
