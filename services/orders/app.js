'use strict';

const express = require('express');
const { OrderBook } = require('./orders');

function buildApp(book = new OrderBook()) {
  const app = express();
  app.use(express.json());

  app.get('/health', (req, res) => {
    res.json({ status: 'ok', service: 'orders', orders: book.all().length });
  });

  app.get('/orders', (req, res) => {
    res.json({ orders: book.all(req.query.status) });
  });

  app.get('/orders/:id', (req, res) => {
    const order = book.find(req.params.id);
    if (!order) {
      return res.status(404).json({ error: 'no such order' });
    }
    return res.json(order);
  });

  app.post('/orders', (req, res) => {
    try {
      return res.status(201).json(book.place(req.body || {}));
    } catch (error) {
      return res.status(400).json({ error: error.message });
    }
  });

  app.post('/orders/:id/ship', (req, res) => {
    const order = book.ship(req.params.id);
    if (!order) {
      return res.status(409).json({ error: 'that order cannot be shipped' });
    }
    return res.json(order);
  });

  app.post('/orders/:id/cancel', (req, res) => {
    const order = book.cancel(req.params.id);
    if (!order) {
      return res.status(409).json({ error: 'that order cannot be cancelled' });
    }
    return res.json(order);
  });

  return app;
}

module.exports = { buildApp };
