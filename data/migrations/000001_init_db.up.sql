--create users table
CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "email" TEXT NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS "email_users_idx" ON "users"("email");

CREATE TYPE "relationship_type" AS ENUM ('friend', 'subscribe', 'block');

--create relationships table
CREATE TABLE "relationships" (
  "id" SERIAL PRIMARY KEY,
  "requester_id" INT NOT NULL,
  "target_id" INT NOT NULL,
  "type" relationship_type NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  FOREIGN KEY ("requester_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY ("target_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "requester_id_target_id_relationships_idx" ON "relationships"("requester_id", "target_id","type");
