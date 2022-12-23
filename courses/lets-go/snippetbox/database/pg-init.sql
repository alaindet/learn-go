/* Table: snippets */
CREATE TABLE "snippets" (
  "id" serial NOT NULL,
  PRIMARY KEY ("id"),
  "title" character varying(100) NOT NULL,
  "content" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "expires_at" timestamp NOT NULL
);

ALTER TABLE "snippets"
DROP CONSTRAINT "snippets_pkey";
CREATE INDEX "snippets_created_at" ON "snippets" ("created_at");

/* Table: sessions */
CREATE TABLE "sessions" (
	token TEXT PRIMARY KEY,
	data BYTEA NOT NULL,
	expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

/* Table: users */
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

/* Data */
INSERT INTO
  snippets (title, content, created_at, expires_at)
VALUES
  (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '365' DAY
  ),
  (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '365' DAY
  ),
  (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '7' DAY
  );
