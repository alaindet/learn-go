-- Users table
DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "users" (
  "id" integer DEFAULT nextval('users_id_seq') NOT NULL,
  "name" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "password" character varying(255) NOT NULL,
  "created_at" timestamptz NOT NULL,
  CONSTRAINT "users_email" UNIQUE ("email"),
  CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

-- Users data
INSERT INTO "users" ("name", "email", "password", "created_at")
VALUES (
  'Mock User',
  'mock@example.com',
  -- Password is "mock@example.com", it's bcrypt-hashed with 12 rounds
  '$2a$12$s8rnYDY559CU1Rb1VbcBU.eEo2tFCGxoXiQEVri1hwJkmRuhyYejy',
  now()
);

-- Sessions table
DROP TABLE IF EXISTS "sessions";
CREATE TABLE "sessions" (
  "token" text NOT NULL,
  "data" bytea NOT NULL,
  "expiry" timestamptz NOT NULL,
  CONSTRAINT "sessions_pkey" PRIMARY KEY ("token")
) WITH (oids = false);

CREATE INDEX "sessions_expiry_idx" ON "sessions" USING btree ("expiry");

-- Snippets table
DROP TABLE IF EXISTS "snippets";
DROP SEQUENCE IF EXISTS snippets_id_seq;
CREATE SEQUENCE snippets_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "snippets" (
  "id" integer DEFAULT nextval('snippets_id_seq') NOT NULL,
  "title" character varying(100) NOT NULL,
  "content" text NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "expires_at" timestamp NOT NULL
) WITH (oids = false);

CREATE INDEX "snippets_created_at" ON "snippets" USING btree ("created_at");
