'use strict';

const request = require('supertest');
const { buildApp } = require('../app');

describe('the orders API', () => {
  let app;
  const order = { lines: [{ sku: 'SKU-1', quantity: 2, price: 5 }], total: 10 };

  beforeEach(() => {
    app = buildApp();
  });

  test('GET /health is ok', async () => {
    const response = await request(app).get('/health');
    expect(response.status).toBe(200);
    expect(response.body.service).toBe('orders');
  });

  test('POST /orders creates one', async () => {
    const response = await request(app).post('/orders').send(order);
    expect(response.status).toBe(201);
    expect(response.body.total).toBe(10);
  });

  test('POST /orders rejects an empty order', async () => {
    expect((await request(app).post('/orders').send({})).status).toBe(400);
  });

  test('an unknown order is a 404', async () => {
    expect((await request(app).get('/orders/999')).status).toBe(404);
  });

  test('shipping twice is a conflict', async () => {
    const created = await request(app).post('/orders').send(order);
    const id = created.body.id;

    expect((await request(app).post(`/orders/${id}/ship`)).status).toBe(200);
    expect((await request(app).post(`/orders/${id}/ship`)).status).toBe(409);
  });
});
