import Link from "next/link";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Typography,
  Button,
} from "@mui/material";
import axios from "axios";
import { headers } from "next/headers";

export default async function Home() {
  const res = await axios.get(`${process.env.API_GATEWAY_URL}/api/tickets`, {
    headers: {
      ...Object.fromEntries(headers().entries()),
    },
  });
  const tickets = res.data;

  return (
    <>
      <Typography variant="h4" gutterBottom>
        Tickets
      </Typography>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Title</TableCell>
              <TableCell>Price</TableCell>
              <TableCell>Link</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {
              // eslint-disable-next-line @typescript-eslint/no-explicit-any
              tickets.map((ticket: any) => (
                <TableRow key={ticket.id}>
                  <TableCell>{ticket.title}</TableCell>
                  <TableCell>{ticket.price}</TableCell>
                  <TableCell>
                    <Link href={`/tickets/${ticket.id}`} passHref>
                      <Button variant="contained" color="primary">
                        View
                      </Button>
                    </Link>
                  </TableCell>
                </TableRow>
              ))
            }
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
