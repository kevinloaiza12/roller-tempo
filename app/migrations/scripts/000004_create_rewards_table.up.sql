CREATE TABLE IF NOT EXISTS rewards (
    RewardName             VARCHAR(255) PRIMARY KEY,
    RewardDescription      VARCHAR(255) NOT NULL,
    RewardPrice            INT NOT NULL
);