CREATE TYPE "relationship_type" AS ENUM ('friend', 'subscribe', 'block');

--create relationships table
CREATE TABLE "relationships" (
   "id" SERIAL PRIMARY KEY,
   "requester_email" VARCHAR NOT NULL,
   "target_email" VARCHAR NOT NULL,
   "type" relationship_type NOT NULL,
   "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
   "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX "requester_email_target_email_type_relationships_idx" ON "relationships"("requester_email", "target_email","type");
