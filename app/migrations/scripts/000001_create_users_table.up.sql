CREATE TABLE IF NOT EXISTS usuarios (
  ID                 SERIAL INT PRIMARY KEY,
  Turno              INT NULL,
  Monedas            INT DEFAULT 0
);