CREATE TABLE IF NOT EXISTS "branch" (
  "id" VARCHAR PRIMARY KEY,
  "branch_code" VARCHAR(10) NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "address" VARCHAR(255) NOT NULL,
  "phone_number" VARCHAR(13) NOT NULL
);