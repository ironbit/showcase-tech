#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$DATABASE_USER" --dbname "$DATABASE_NAME" <<-EOSQL
  CREATE TABLE IF NOT EXISTS "author" (
    "author_id" serial NOT NULL PRIMARY KEY,
    "author_name" varchar(128) NOT NULL DEFAULT '' 
  );

  CREATE TABLE IF NOT EXISTS "book" (
    "book_id" serial NOT NULL PRIMARY KEY,
    "book_name" varchar(128) NOT NULL DEFAULT '' ,
    "book_genre" varchar(64) NOT NULL DEFAULT '' 
  );

  CREATE TABLE IF NOT EXISTS "book_authors" (
    "id" serial NOT NULL PRIMARY KEY,
    "book_id" integer NOT NULL,
    "author_id" integer NOT NULL
  );

  CREATE SEQUENCE IF NOT EXISTS author_seq
  START WITH 1
  OWNED BY author.author_id;

  CREATE SEQUENCE IF NOT EXISTS book_seq
  START WITH 1
  OWNED BY book.book_id;

  CREATE SEQUENCE IF NOT EXISTS book_authors_seq
  START WITH 1
  OWNED BY book_authors.id;
EOSQL
