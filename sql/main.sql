CREATE TABLE vendas (
    id SERIAL PRIMARY KEY,
    data_venda TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE itens_venda (
    id SERIAL PRIMARY KEY,
    venda_id INTEGER REFERENCES vendas(id) ON DELETE CASCADE,
    codigo_interno VARCHAR(50) NOT NULL,
    codigo_barras VARCHAR(50) NOT NULL,
    quantidade_vendida INTEGER NOT NULL
);