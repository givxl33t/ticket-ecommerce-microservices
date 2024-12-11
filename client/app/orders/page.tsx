import { Container, Typography, Card, CardContent } from "@mui/material";
import Grid from "@mui/material/Grid2"; // Import Grid2
import axios from "axios";
import { headers } from "next/headers";

export interface IOrder {
  id: string;
  version: number;
  ticket: {
    title: string;
    price: number;
  };
  status: string;
  expiresAt: string;
}

export default async function Orders() {
  const res = await axios.get(`${process.env.API_GATEWAY_URL}/api/orders`, {
    headers: {
      ...Object.fromEntries(headers().entries()),
    },
  });
  const orders = res.data;

  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Typography variant="h4" component="h1" textAlign="center" gutterBottom>
        My Orders
      </Typography>

      {orders.length === 0 ? (
        <Typography variant="body1" color="textSecondary" textAlign="center">
          You have no orders at the moment.
        </Typography>
      ) : (
        <Grid container spacing={3}>
          {orders.map((order: IOrder) => (
            <Grid size={{ xs: 12, sm: 6, md: 4 }} key={order.id}>
              <Card variant="outlined" sx={{ height: "100%" }}>
                <CardContent>
                  <Typography variant="h6" component="div">
                    {order.ticket.title}
                  </Typography>
                  <Typography variant="body2" color="textSecondary">
                    Status: {order.status}
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      )}
    </Container>
  );
}
