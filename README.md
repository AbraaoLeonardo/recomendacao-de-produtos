# Recomendacao (Go + PostgreSQL)

Small Go service that stores sales and inventory in PostgreSQL and returns product recommendations via HTTP.

## Features
- Persist sales records in PostgreSQL
- Simple recommendation endpoint that returns related products
- Run locally with a single command

## Requirements
- Docker & Docker Compose (or a compatible environment)
- make
- Go (for development)

## Quick start
1. Open a terminal in the project folder.
2. Run:
```bash
make up
```
This starts the API and a PostgreSQL database (via Docker Compose).

## API (examples)
Assumed base URL: http://localhost:8080

- Record a sale
    - POST /sales
    - Body (application/json):
        ```json
            {
            "data_venda": "2025-11-02 11:11:54",
            "produtos": [
                {
                "codigo_interno": "001",
                "codigo_barras": "7891234567890",
                "quantidade_vendida": 2
                },
                {
                "codigo_interno": "002",
                "codigo_barras": "7890987654321",
                "quantidade_vendida": 1
                },
                {
                "codigo_interno": "003",
                "codigo_barras": "7891122334455",
                "quantidade_vendida": 5
                }
            ]
            }
        ```
    - Response: 201 Created

- Get recommendation
    - POST /recommend
    - Body:
        ```json
        {
            "product_id": "sku-123",
            "limit": 5
        }
        ```
    - Response (example):
        ```json
            {
                "produto_buscado": sku-123,
                "data_limite":     '2026-01-01', //data inicial considerada para analise
                "associados":      associados, // lista de produtos associados
            }
        ```

## Development notes
- Service implemented in Go; configuration via environment variables (DB connection, port).


## Work In Progress
- [ ] Create the terraform to deploy on Google Cloud Plataform
- [ ] Return a list of products
- [ ] Create an authentication process
- [ ] Make the docker-compose to run local.