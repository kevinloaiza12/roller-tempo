CREATE TABLE IF NOT EXISTS users (
  UserID                 INT NOT NULL PRIMARY KEY,
  UserCoins              INT DEFAULT 0,
  UserTurn               INT DEFAULT 0
);