CREATE TABLE IF NOT EXISTS atracciones (
  Nombre             VARCHAR(255) PRIMARY KEY,
  Descripcion        VARCHAR(255) DEFAULT '',
  Duracion           INT NOT NULL,
  Capacidad          INT NOT NULL,
  Turno_Actual    INT DEFAULT 1,
  Siguiente_Turno    INT DEFAULT 1
);
