"use client";

import { useEffect, useState } from "react";
import { loadStripe } from "@stripe/stripe-js";
import { Elements } from "@stripe/react-stripe-js";
import { Box, Container, Paper, Typography } from "@mui/material";
import { IOrder } from "../page";
import PaymentForm from "./PaymentForm";

const stripePromise = loadStripe(
  "pk_test_51NuxwkEPFj7I6FUPEPgKjkz5nGKUa9fgZ5QrBl28z5xSn2yYR3tbwOCwekEiguJRp6OttXtxsqO2CMnOUblPY8Xw00RN1EBRDt",
);

export default function OrderShow({ order }: { order: IOrder }) {
  const [timeLeft, setTimeLeft] = useState(0);

  useEffect(() => {
    const findTimeLeft = () => {
      const msLeft = new Date(order.expiresAt).getTime() - new Date().getTime();
      setTimeLeft(Math.round(msLeft / 1000));
    };

    findTimeLeft();
    const timerId = setInterval(findTimeLeft, 1000);

    return () => clearInterval(timerId);
  }, [order]);

  if (timeLeft < 0) {
    return (
      <Container maxWidth="sm" sx={{ mt: 4, textAlign: "center" }}>
        <Typography variant="h5" color="error">
          Order expired
        </Typography>
      </Container>
    );
  }

  return (
    <Container maxWidth="sm" sx={{ mt: 4 }}>
      <Paper sx={{ p: 4 }}>
        <Box sx={{ textAlign: "center", mb: 4 }}>
          <Typography variant="h5">Order Details</Typography>
          <Typography variant="body1" sx={{ mt: 2 }}>
            Time left to pay: <strong>{timeLeft}</strong> seconds
          </Typography>
          <Typography variant="body2" sx={{ mt: 1 }}>
            Ticket: {order.ticket.title}
          </Typography>
          <Typography variant="body2" sx={{ mb: 2 }}>
            Price: ${order.ticket.price}
          </Typography>
        </Box>
        <Elements stripe={stripePromise}>
          <PaymentForm order={order} />
        </Elements>
      </Paper>


    </Container>
  );
}
