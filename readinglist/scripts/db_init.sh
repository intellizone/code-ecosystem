#!/usr/bin/env bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE DATABASE readinglist;
	CREATE USER readinglist WITH PASSWORD 'passW09d';
	GRANT ALL PRIVILEGES ON DATABASE readinglist TO readinglist;
	
	-- Disable remote postgres superuser login
	ALTER USER postgres WITH PASSWORD NULL;
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "readinglist" <<-EOSQL
	CREATE TABLE IF NOT EXISTS books (
	    id bigserial PRIMARY KEY,
	    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	    title text NOT NULL,
	    published integer NOT NULL,
	    pages integer NOT NULL,
	    genere text[] NOT NULL,
	    rating real NOT NULL,
	    version integer NOT NULL DEFAULT 1
	);
	
	GRANT SELECT, INSERT, UPDATE, DELETE ON books TO readinglist;
	GRANT USAGE, SELECT ON SEQUENCE books_id_seq TO readinglist;
EOSQL
