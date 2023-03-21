CREATE TABLE IF NOT EXISTS atracciones (
  Nombre             VARCHAR(255) PRIMARY KEY,
  Descripcion        VARCHAR(255) DEFAULT '',
  Duracion           INT NOT NULL,
  Capacidad          INT NOT NULL,
  Siguiente_Turno    INT DEFAULT 1
);