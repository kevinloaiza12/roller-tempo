CREATE TABLE IF NOT EXISTS attractions (
  AttractionName               VARCHAR(255) PRIMARY KEY,
  AttractionDescription        VARCHAR(255) DEFAULT '',
  AttractionCapacity           INT NOT NULL,
  AttractionDuration           INT NOT NULL,
  AttractionCurrentTurn        INT DEFAULT 1,
  AttractionNextTurn           INT DEFAULT 1,
  AttractionPosX               FLOAT NOT NULL,
  AttractionPosY               FLOAT NOT NULL
);