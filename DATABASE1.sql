-- @conn=tic2
CREATE DATABASE tic2;



CREATE TABLE games (
    id              UUID PRIMARY KEY,
    board           JSONB NOT NULL,
    current_player  INT NOT NULL DEFAULT 1,
    status          TEXT NOT NULL DEFAULT 'active',
    winner          INT
);