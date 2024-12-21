"use client";

import { Box, Button, Container, TextField, Typography } from "@mui/material";
import { useRouter } from "next/navigation";
import { FormEvent, useState } from "react";
import useRequest from "@/hooks/useRequest";
import { useAuth } from "@/providers/auth";

export default function SignIn() {
  const { refreshUser } = useAuth();
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { doRequest, errors } = useRequest({
    url: "/api/users/signin",
    method: "post",
    body: {
      email,
      password,
    },
    onSuccess: () => {
      refreshUser();
      router.push("/");
    },
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
          Sign In
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
          Sign In
        </Button>
      </Box>
    </Container>
  );
}
