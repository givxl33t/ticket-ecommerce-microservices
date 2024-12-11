"use client"

import { useStripe, useElements, CardElement } from "@stripe/react-stripe-js";
import { IOrder } from "../page";
import { FormEvent, useState } from "react";
import useRequest from "@/hooks/useRequest";
import { useRouter } from "next/navigation";
import { useAuth } from "@/providers/auth";
import { Box, Button, CircularProgress } from "@mui/material";

export default function PaymentForm({ order }: { order: IOrder }) {
  const { currentUser } = useAuth();
  const stripe = useStripe()
  const elements = useElements()
  const router = useRouter()
  const [loading, setLoading] = useState(false);
  const { doRequest, errors } = useRequest({
    url: "/api/payments",
    method: "post",
    body: {
      orderId: order.id,
    },
    onSuccess: () => router.push("/orders"),
  });

  const onSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setLoading(true);

    if (!stripe || !elements) {
      console.error("Stripe or elements is not initialized");
      return;
    }

    const cardElement = elements.getElement(CardElement);
    if (!cardElement) {
      console.error("Card element is not initialized");
      return;
    }

    const { id: clientSecretId } = await doRequest();
    const { error } = await stripe.confirmCardPayment(clientSecretId, {
      payment_method: {
        card: cardElement,
        billing_details: {
          email: currentUser!.email,
        }
      }
    })

    if (error) {
      console.error("Payment Failed", error);
    }

    setLoading(false);
  };

  return (
    <form onSubmit={onSubmit}>
      <Box sx={{ mb: 3 }}>
        <CardElement
          options={{
            style: {
              base: {
                fontSize: "16px",
                color: "#424770",
                "::placeholder": { color: "#aab7c4" },
              },
              invalid: { color: "#9e2146" },
            },
          }}
        />
      </Box>

      {errors && <Box sx={{ mb: 2 }}>{errors}</Box>}

      <Button
        type="submit"
        variant="contained"
        color="primary"
        fullWidth
        disabled={!stripe || loading}
      >
        {loading ? <CircularProgress size={24} color="inherit" /> : "Pay Now"}
      </Button>
    </form>
  );
}