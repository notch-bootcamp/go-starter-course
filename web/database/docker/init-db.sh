#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER docker WITH ENCRYPTED PASSWORD 'mypassword';
	CREATE DATABASE docker;
	GRANT ALL PRIVILEGES ON DATABASE docker TO docker;
	GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO docker;
	ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO docker;
  \c docker;
	CREATE TABLE position (
  	position_id serial PRIMARY KEY,
  	label VARCHAR ( 50 ) UNIQUE NOT NULL,
  	latitude FLOAT8 NOT NULL,
    longitude FLOAT8 NOT NULL,
    elevation INT
  );
  ALTER TABLE position owner TO docker;
EOSQL
