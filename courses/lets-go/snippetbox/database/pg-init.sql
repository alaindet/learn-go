/* Tables */
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
