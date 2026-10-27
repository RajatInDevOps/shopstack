'use strict';

// Order bookkeeping, with no HTTP in it.

class OrderBook {
  constructor() {
    this.orders = [];
    this.nextId = 1;
  }

  place({ lines, total }) {
    if (!Array.isArray(lines) || lines.length === 0) {
      throw new Error('an order needs at least one line');
    }
    const order = {
      id: this.nextId++,
      lines,
      total: Number(total) || 0,
      status: 'placed',
      placed_at: new Date().toISOString(),
    };
    this.orders.push(order);
    return order;
  }

  find(id) {
    return this.orders.find((order) => order.id === Number(id));
  }

  all(status) {
    return status ? this.orders.filter((order) => order.status === status) : this.orders;
  }

  // An order can be cancelled right up until it ships.
  cancel(id) {
    const order = this.find(id);
    if (!order || order.status === 'shipped') {
      return null;
    }
    order.status = 'cancelled';
    return order;
  }

  ship(id) {
    const order = this.find(id);
    if (!order || order.status !== 'placed') {
      return null;
    }
    order.status = 'shipped';
    return order;
  }
}

module.exports = { OrderBook };
