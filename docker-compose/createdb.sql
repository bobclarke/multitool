CREATE DATABASE bird_encyclopedia;
\connect bird_encyclopedia

CREATE TABLE birds (
  id SERIAL PRIMARY KEY,
  bird VARCHAR(256),
  description VARCHAR(1024)
);