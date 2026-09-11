-- @conn=tic2
CREATE DATABASE tic2;



CREATE TABLE games (
    id              UUID PRIMARY KEY,
    board           JSONB NOT NULL,
    current_player  INT NOT NULL DEFAULT 1,
    status          TEXT NOT NULL DEFAULT 'active',
    winner          INT
);


CREATE TABLE users (
    id UUID PRIMARY KEY,
    login TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

DROP TABLE IF EXISTS games;

CREATE TABLE games (
    id           UUID PRIMARY KEY,
    board        JSONB NOT NULL,
    player_x     UUID REFERENCES users(id),
    player_o     UUID REFERENCES users(id),
    state        TEXT NOT NULL,
    turn_player  UUID,
    winner       UUID,
    vs_computer  BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS idx_games_state ON games (state);