"""The gateway — the one service a browser talks to."""

import os

import httpx
from fastapi import FastAPI, HTTPException

from basket import Basket

ORDERS_URL = os.environ.get("ORDERS_URL", "http://orders:3000")
INVENTORY_URL = os.environ.get("INVENTORY_URL", "http://inventory:8080")

app = FastAPI(title="ShopStack Gateway")
basket = Basket()


@app.get("/health")
def health():
    return {"status": "ok", "service": "gateway", "basket_items": basket.count()}


@app.get("/basket")
def read_basket():
    return {"items": basket.items(), "total": basket.total()}


@app.post("/basket/{sku}")
def add_to_basket(sku: str, quantity: int = 1, price: float = 0.0):
    if quantity < 1:
        raise HTTPException(status_code=400, detail="quantity must be at least 1")
    basket.add(sku, quantity, price)
    return {"items": basket.items(), "total": basket.total()}


@app.delete("/basket/{sku}")
def remove_from_basket(sku: str):
    if not basket.remove(sku):
        raise HTTPException(status_code=404, detail="that sku is not in the basket")
    return {"items": basket.items(), "total": basket.total()}


@app.get("/catalogue")
async def catalogue():
    """Reads through to the inventory service."""
    async with httpx.AsyncClient(timeout=5.0) as client:
        response = await client.get(INVENTORY_URL + "/items")
    return response.json()


@app.post("/checkout")
async def checkout():
    if basket.count() == 0:
        raise HTTPException(status_code=400, detail="the basket is empty")

    payload = {"lines": basket.items(), "total": basket.total()}
    async with httpx.AsyncClient(timeout=5.0) as client:
        response = await client.post(ORDERS_URL + "/orders", json=payload)

    basket.clear()
    return response.json()
