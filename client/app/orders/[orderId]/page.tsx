import OrderShow from "./OrderShow";
import axios from "axios";
import { headers } from "next/headers";
import { IOrder } from "../page";

export default async function Page({ params }: { params: { orderId: string } }) {
  const { orderId } = params;
  const order: IOrder = await axios
    .get(`${process.env.API_GATEWAY_URL}/api/orders/${orderId}`, {
      headers: {
        ...Object.fromEntries(headers().entries()),
      },
    })
    .then((res) => res.data);

  return <OrderShow order={{ ...order }} />;
}
