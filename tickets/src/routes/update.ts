import express, { Request, Response } from 'express';
import { body } from 'express-validator';
import {
  validateRequest,
  NotFoundError,
  requireAuth,
  NotAuthorizedError,
  BadRequestError,
} from '@romen-tix-micro/common';
import { Ticket } from '../models/ticket';
import { TicketUpdatedPublisher } from '../events/publishers/ticket-updated-publisher';
import { natsWrapper } from '../nats-wrapper';

const router = express.Router();

router.put(
  '/api/tickets/:id',
  requireAuth,
  [
    body('title').not().isEmpty().withMessage('Title is required.'),
    body('price').isFloat({ gt: 0 }).withMessage('Price must be greater than 0.'),
  ],
  validateRequest,
  async (req: Request, res: Response) => {
    // Find the ticket the user is trying to edit
    const ticket = await Ticket.findById(req.params.id);

    // If no ticket, throw error
    if (!ticket) {
      throw new NotFoundError();
    }

    // If ticket is reserved, throw error
    if (ticket.orderId) {
      throw new BadRequestError('Cannot edit a reserved ticket.');
    }

    // If ticket does not belong to user, throw error
    if (ticket.userId !== req.currentUser!.id) {
      throw new NotAuthorizedError();
    }

    // Update the ticket
    ticket.set({
      title: req.body.title,
      price: req.body.price,
    });

    // Save the ticket
    await ticket.save();
    await new TicketUpdatedPublisher(natsWrapper.client).publish({
      id: ticket.id,
      title: ticket.title,
      price: ticket.price,
      userId: ticket.userId,
      version: ticket.version,
    });

    // Send the ticket
    res.send(ticket);
  }
);

export { router as updateTicketRouter };