"use client";

import { useState } from "react";
import { Box, Button, Container, TextField, Typography } from "@mui/material";
import { useRouter } from "next/navigation";
import useRequest from "@/hooks/useRequest";

export default function NewTicket() {
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState("");
  const router = useRouter();
  const { doRequest, errors } = useRequest({
    url: "/api/tickets",
    method: "post",
    body: { title, price },
    onSuccess: () => {
      router.push("/");
      router.refresh();
    }
  });

  const onSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    doRequest();
  };

  const onBlur = () => {
    const value = parseFloat(price);
    if (isNaN(value)) {
      return;
    }
    setPrice(value.toFixed(2));
  };

  return (
    <Container maxWidth="sm">
      <Box
        component="form"
        onSubmit={onSubmit}
        sx={{
          mt: 4,
          p: 3,
          borderRadius: 2,
          boxShadow: 3,
          bgcolor: "background.paper",
          display: "flex",
          flexDirection: "column",
          gap: 3,
        }}
      >
        <Typography variant="h4" component="h1" textAlign="center">
          Create a Ticket
        </Typography>

        <TextField
          label="Title"
          variant="outlined"
          fullWidth
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />

        <TextField
          label="Price"
          variant="outlined"
          fullWidth
          value={price}
          onChange={(e) => setPrice(e.target.value)}
          onBlur={onBlur}
          required
        />

        {errors && <Box sx={{ width: "100%", mt: 2 }}>{errors}</Box>}

        <Button type="submit" variant="contained" color="primary" size="large" fullWidth>
          Submit
        </Button>
      </Box>
    </Container>
  );
}