CREATE TABLE IF NOT EXISTS atracciones (
  ID                 INT PRIMARY KEY,
  Nombre             VARCHAR(255) NOT NULL,
  Descripcion        VARCHAR(255) NULL,
  Duracion           INT NOT NULL,
  Capacidad          INT DEFAULT 1,
  Siguiente_Turno    INT DEFAULT 1
);