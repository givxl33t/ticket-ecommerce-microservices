import express, { Request, Response } from 'express';
import { body } from 'express-validator';
import { 
  requireAuth, 
  validateRequest,
  BadRequestError,
  NotFoundError,
  NotAuthorizedError,
  OrderStatus,
} from '@romen-tix-micro/common';
import { Order } from '../models/order';
import { stripe } from '../stripe'; 
import { Payment } from '../models/payment';
import { PaymentCreatedPublisher } from '../events/publishers/payment-created-publisher';
import { natsWrapper } from '../nats-wrapper';

const router = express.Router();

router.post(
  '/api/payments',
  requireAuth,
  [
    body('orderId')
      .not()
      .isEmpty()
      .withMessage('Order ID is required'),
  ],
  validateRequest,
  async (req: Request, res: Response) => {
    const { orderId } = req.body;

    const order = await Order.findById(orderId);

    if (!order) {
      throw new NotFoundError();
    }

    if (order.userId !== req.currentUser!.id) {
      throw new NotAuthorizedError();
    }

    if (order.status === OrderStatus.Cancelled) {
      throw new BadRequestError('Cannot pay for a cancelled order');
    }

    const paymentIntent = await stripe.paymentIntents.create({
      amount: order.price * 100,
      currency: 'usd',
      description: `Ticket for ${order.id}`,
      payment_method_types: ['card'],
      metadata: {
        orderId: order.id,
      }
    });

    const payment = Payment.build({
      orderId,
      stripeId: paymentIntent.client_secret as string,
    });
    await payment.save();

    new PaymentCreatedPublisher(natsWrapper.client).publish({
      id: payment.id,
      orderId: payment.orderId,
      stripeId: payment.stripeId,
    });

    res.status(201).send({ id: payment.stripeId });
  }
)

export { router as createChargeRouter };