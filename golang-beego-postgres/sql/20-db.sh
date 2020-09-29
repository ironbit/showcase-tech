#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$DATABASE_NAME" <<-EOSQL
  CREATE TABLE IF NOT EXISTS "author" (
    "author_id" integer NOT NULL PRIMARY KEY,
    "author_name" varchar(128) NOT NULL DEFAULT '' 
  );

  CREATE TABLE IF NOT EXISTS "book" (
    "book_id" integer NOT NULL PRIMARY KEY,
    "book_name" varchar(128) NOT NULL DEFAULT '' ,
    "book_genre" varchar(64) NOT NULL DEFAULT '' 
  );

  CREATE TABLE IF NOT EXISTS "book_authors" (
    "id" serial NOT NULL PRIMARY KEY,
    "book_id" integer NOT NULL,
    "author_id" integer NOT NULL
  );
EOSQL
