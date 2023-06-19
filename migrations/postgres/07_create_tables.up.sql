CREATE TABLE "supplier" (
  "id" VARCHAR NOT NULL PRIMARY KEY,
  "first_name" VARCHAR(255) NOT NULL,
  "last_name" VARCHAR(255) NOT NULL,  
  "store_id" VARCHAR(36) NOT NULL REFERENCES "store"("id"),
  "phone_number" VARCHAR(13) NOT NULL,
  "active" VARCHAR(100) NOT NULL
);