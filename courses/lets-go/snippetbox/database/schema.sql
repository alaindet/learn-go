/* Create a `snippets` table. */

-- MySQL
-- CREATE TABLE snippets (
--   id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
--   title VARCHAR(100) NOT NULL,
--   content TEXT NOT NULL,
--   created DATETIME NOT NULL,
--   expires DATETIME NOT NULL
-- );
--
-- CREATE INDEX idx_snippets_created ON snippets(created);

-- Postgres
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
