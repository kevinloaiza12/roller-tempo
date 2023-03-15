CREATE TABLE IF NOT EXISTS atracciones (
  ID                 SERIAL PRIMARY KEY,
  Nombre             VARCHAR(255) NOT NULL,
  Descripcion        VARCHAR(255) DEFAULT '',
  Duracion           INT DEFAULT 0,
  Capacidad          INT DEFAULT 1,
  Siguiente_Turno    INT DEFAULT 1
);