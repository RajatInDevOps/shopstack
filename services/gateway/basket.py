"""The basket, with no HTTP in it so it can be tested on its own."""


class Basket:
    def __init__(self):
        self._lines = {}

    def add(self, sku, quantity=1, price=0.0):
        line = self._lines.setdefault(sku, {"sku": sku, "quantity": 0, "price": price})
        line["quantity"] += quantity
        if price:
            line["price"] = price
        return line

    def remove(self, sku):
        return self._lines.pop(sku, None) is not None

    def items(self):
        return [dict(line) for line in self._lines.values()]

    def count(self):
        return sum(line["quantity"] for line in self._lines.values())

    def total(self):
        return round(sum(line["quantity"] * line["price"] for line in self._lines.values()), 2)

    def clear(self):
        self._lines.clear()
