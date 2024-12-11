"use client";

import { Box, Button, Container, TextField, Typography } from "@mui/material";
import useRequest from "@/hooks/useRequest";
import { FormEvent, useState } from "react";
import { useRouter } from "next/navigation";

export default function SignUp() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();
  const { doRequest, errors } = useRequest({
    url: "/api/users/signup",
    method: "post",
    body: {
      email,
      password,
    },
    onSuccess: () => router.push("/"),
  });

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    await doRequest();
  };

  return (
    <Container maxWidth="xs">
      <Box
        component="form"
        onSubmit={handleSubmit}
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          padding: 3,
          borderRadius: 1,
          boxShadow: 3,
          backgroundColor: "white",
        }}
      >
        <Typography variant="h5" sx={{ marginBottom: 2 }}>
          Sign Up
        </Typography>

        <TextField
          name="email"
          label="Email"
          type="email"
          onChange={(e) => setEmail(e.target.value)}
          variant="outlined"
          fullWidth
          required
          sx={{ marginBottom: 2 }}
        />

        <TextField
          name="password"
          label="Password"
          type="password"
          onChange={(e) => setPassword(e.target.value)}
          variant="outlined"
          fullWidth
          required
          sx={{ marginBottom: 2 }}
        />

        {errors && <Box sx={{ width: "100%", marginBottom: "1rem" }}>{errors}</Box>}

        <Button
          type="submit"
          variant="contained"
          color="primary"
          fullWidth
          sx={{ marginBottom: 2 }}
        >
          Sign Up
        </Button>
      </Box>
    </Container>
  );
}
