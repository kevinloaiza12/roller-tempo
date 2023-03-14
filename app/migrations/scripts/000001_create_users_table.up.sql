CREATE TABLE IF NOT EXISTS users (
  ID                 INT PRIMARY KEY,
  Turno              INT NULL,
  Cantidad_Monedas   INT DEFAULT 0
);