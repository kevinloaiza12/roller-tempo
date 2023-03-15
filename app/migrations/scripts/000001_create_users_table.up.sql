CREATE TABLE IF NOT EXISTS usuarios (
  ID                 SERIAL PRIMARY KEY,
  Turno              INT DEFAULT 0,
  Monedas            INT DEFAULT 0
);