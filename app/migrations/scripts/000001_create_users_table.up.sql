CREATE TABLE IF NOT EXISTS users (
  UserID                 INT NOT NULL PRIMARY KEY,
  UserTurn               INT DEFAULT 0,
  UserCoins              INT DEFAULT 0
);