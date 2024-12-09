'use client'

import { fetchApi } from "@/lib/apiWrapper";
import { Box, Button, Container, TextField, Typography } from "@mui/material";
import { useRouter } from "next/navigation";
import { FormEvent } from "react";

const SignUp = () => {
  const router = useRouter()

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    const formData = new FormData(event.currentTarget)

    try {
      const email = formData.get('email') as string
      const password = formData.get('password') as string
      await fetchApi('/api/users/signup', {
        method: 'POST',
        body: JSON.stringify({
          email, password
        })
      })
      router.push('/')
    } catch (error) {
      console.error("Failed", error)
    }
  }

  return (
    <Container maxWidth="xs">
      <Box
        component="form"
        onSubmit={handleSubmit}
        sx={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          padding: 3,
          borderRadius: 1,
          boxShadow: 3,
          backgroundColor: 'white',
          mt: 15
        }}
      >
        <Typography variant="h5" sx={{ marginBottom: 2 }}>
          Sign Up
        </Typography>

        <TextField
          name="email"
          label="Email"
          variant="outlined"
          fullWidth
          required
          sx={{ marginBottom: 2 }}
        />
        
        <TextField
          name="password"
          label="Password"
          type="password"
          variant="outlined"
          fullWidth
          required
          sx={{ marginBottom: 2 }}
        />

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
};

export default SignUp