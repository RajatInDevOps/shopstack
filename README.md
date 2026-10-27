# ShopStack

A small e-commerce stack, split the way a real one is: three services in three
languages, with a database and a cache behind them.

| Service     | Language          | Port | What it owns                    |
|-------------|-------------------|------|---------------------------------|
| `gateway`   | Python / FastAPI  | 8000 | the basket, and the one public API |
| `orders`    | Node / Express    | 3000 | orders and their lifecycle      |
| `inventory` | Go                | 8080 | warehouse stock                 |
| `postgres`  | Postgres 16       | 5432 | order storage                   |
| `redis`     | Redis 7           | 6379 | session cache                   |

The gateway is the only service a browser talks to; it reads through to
`inventory` for the catalogue and posts to `orders` at checkout.

## Running the whole stack

```bash
docker compose up --build
curl localhost:18000/health
curl localhost:18000/catalogue
```

The host ports are 18000 (gateway), 13000 (orders) and 18080 (inventory);
inside the network the services still talk to each other on 8000, 3000 and
8080.

## Tests

```bash
cd services/gateway   && pytest
cd services/orders    && npm test
cd services/inventory && go test ./...
```
