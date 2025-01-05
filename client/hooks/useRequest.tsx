/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import axios from "axios";
import { ReactElement, useState } from "react";
import { Alert, AlertTitle, List, ListItem } from "@mui/material";

interface IUseRequestProps {
  url: string;
  method: string;
  body?: object;
  onSuccess?: (data: object) => void;
}

const useRequest = ({ url, method, body, onSuccess }: IUseRequestProps) => {
  const [errors, setErrors] = useState<ReactElement | null>(null);
  // If we were to log body here, remember we are in client side
  // so we need to check the logs on the browser on the server side
  const doRequest = async (props = {}) => {
    try {
      let response;
      switch (method) {
        case "get":
          response = await axios.get(url, { ...props });
          break;
        case "post":
          response = await axios.post(url, { ...body, ...props });
          break;
        case "put":
          response = await axios.put(url, { ...body, ...props });
          break;
        case "delete":
          response = await axios.delete(url, { ...props });
          break;
        default:
          throw new Error(`Unsupported method: ${method}`);
      }
      if (onSuccess) {
        onSuccess(response.data);
      }
      setErrors(null);
      return response.data;
    } catch (errors: any) {
      console.log("received errors: ", errors)

      setErrors(
        <Alert severity="error">
          <AlertTitle>Oops...</AlertTitle>
          <List disablePadding>
            {Array.isArray(errors.response.data.message) ? (
              errors.response.data.message.map((err: any) => (
                <ListItem key={err.field}>{err.field}: {err.message}</ListItem>
              ))
            ) : (
              <ListItem>{errors.response.data.message}</ListItem>
            )}
          </List>
        </Alert>,
      );
    }
  };

  return { doRequest, errors };
};

export default useRequest;
