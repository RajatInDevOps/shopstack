'use strict';

const { OrderBook } = require('../orders');

describe('OrderBook', () => {
  let book;
  const line = { lines: [{ sku: 'SKU-1', quantity: 1, price: 10 }], total: 10 };

  beforeEach(() => {
    book = new OrderBook();
  });

  test('starts empty', () => {
    expect(book.all()).toHaveLength(0);
  });

  test('refuses an order with no lines', () => {
    expect(() => book.place({ lines: [] })).toThrow('at least one line');
    expect(() => book.place({})).toThrow('at least one line');
  });

  test('a placed order starts in "placed"', () => {
    expect(book.place(line).status).toBe('placed');
  });

  test('ids increment', () => {
    expect(book.place(line).id).toBe(1);
    expect(book.place(line).id).toBe(2);
  });

  test('an order can be cancelled before it ships', () => {
    const order = book.place(line);
    expect(book.cancel(order.id).status).toBe('cancelled');
  });

  test('a shipped order can no longer be cancelled', () => {
    const order = book.place(line);
    book.ship(order.id);
    expect(book.cancel(order.id)).toBeNull();
  });

  test('an order ships only once', () => {
    const order = book.place(line);
    expect(book.ship(order.id).status).toBe('shipped');
    expect(book.ship(order.id)).toBeNull();
  });

  test('filters by status', () => {
    const first = book.place(line);
    book.place(line);
    book.ship(first.id);

    expect(book.all('shipped')).toHaveLength(1);
    expect(book.all('placed')).toHaveLength(1);
  });
});
