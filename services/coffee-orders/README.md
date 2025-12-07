# coffee-orders

Simple in-memory orders service which validates coffee IDs by calling the catalog service.

Endpoints:
- `GET /orders` - list orders
- `POST /orders` - create an order JSON {coffee_id,quantity}

Set `COFFEE_CATALOG_URL` env var to point to the catalog (default: `http://localhost:8081`).

Runs on port `8082`.
