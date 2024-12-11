"use client";

import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { useAuth } from "@/providers/auth";
import { Box, Typography, CircularProgress } from "@mui/material";
import useRequest from "@/hooks/useRequest";

export default function SignOut() {
  const { refreshUser } = useAuth();
  const router = useRouter();
  const { doRequest } = useRequest({
    url: "/api/users/signout",
    method: "post",
    body: {},
    onSuccess: () => {
      refreshUser();
      router.push("/");
    },
  });

  useEffect(() => {
    doRequest();
  }, [doRequest]);

  return (
    <Box
      display="flex"
      flexDirection="column"
      justifyContent="center"
      alignItems="center"
      height="100vh"
      bgcolor="background.default"
    >
      <CircularProgress color="primary" />
      <Typography variant="h6" sx={{ marginTop: 2 }}>
        Signing out...
      </Typography>
    </Box>
  );
}
