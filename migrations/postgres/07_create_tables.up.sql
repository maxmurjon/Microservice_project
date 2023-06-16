CREATE TABLE "supplier" (
  "id" VARCHAR NOT NULL PRIMARY KEY,
  "fist_name" VARCHAR(255) NOT NULL,
  "last_name" VARCHAR(255) NOT NULL,  
  "store" VARCHAR(10) NOT NULL REFERENCES "store"("id"),
  "phone_number" VARCHAR(13) NOT NULL,
  "active" VARCHAR(100) NOT NULL
);