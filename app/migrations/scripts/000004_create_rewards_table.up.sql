CREATE TABLE premios (
    ID               SERIAL PRIMARY KEY,
    Nombre           VARCHAR(255) NOT NULL,
    Descripcion      VARCHAR(255) NULL,
    Precio           INT DEFAULT 0
);