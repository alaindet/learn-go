\connect "snippetbox";

-- Users table
DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."users" (
    "id" integer DEFAULT nextval('users_id_seq') NOT NULL,
    "name" character varying(255) NOT NULL,
    "email" character varying(255) NOT NULL,
    "password" character varying(255) NOT NULL,
    "created_at" timestamptz NOT NULL,
    CONSTRAINT "users_email" UNIQUE ("email"),
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

-- Sessions table
DROP TABLE IF EXISTS "sessions";
CREATE TABLE "public"."sessions" (
    "token" text NOT NULL,
    "data" bytea NOT NULL,
    "expiry" timestamptz NOT NULL,
    CONSTRAINT "sessions_pkey" PRIMARY KEY ("token")
) WITH (oids = false);

CREATE INDEX "sessions_expiry_idx" ON "public"."sessions" USING btree ("expiry");

-- Snippets table
DROP TABLE IF EXISTS "snippets";
DROP SEQUENCE IF EXISTS snippets_id_seq;
CREATE SEQUENCE snippets_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."snippets" (
    "id" integer DEFAULT nextval('snippets_id_seq') NOT NULL,
    "title" character varying(100) NOT NULL,
    "content" text NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "expires_at" timestamp NOT NULL
) WITH (oids = false);

CREATE INDEX "snippets_created_at" ON "public"."snippets" USING btree ("created_at");

-- Snippets data
INSERT INTO
  "snippets" ("title", "content", "created_at", "expires_at")
VALUES
  (
    "An old silent pond",
    "An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō",
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '365' DAY
  ),
  (
    "Over the wintry forest",
    "Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki",
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '365' DAY
  ),
  (
    "First autumn morning",
    "First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo",
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '7' DAY
  );
