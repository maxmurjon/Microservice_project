CREATE TABLE "store" (
  "id" VARCHAR PRIMARY KEY,
  "branch_id" VARCHAR(36) NOT NULL REFERENCES "branch"("id"),
  "name" VARCHAR(255) NOT NULL
);