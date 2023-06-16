CREATE TABLE "employees" (
  "id" VARCHAR PRIMARY KEY,
  "first_name" VARCHAR(100) NOT NULL,
  "last_name" VARCHAR(100) NOT NULL,
  "store_id" VARCHAR NOT NULL REFERENCES "store"("id"),
  "phone_number" VARCHAR(13) NOT NULL,
  "login" VARCHAR(255) NOT NULL,
  "password_hash" VARCHAR(100) NOT NULL,
  "role" VARCHAR(100) NOT NULL
);