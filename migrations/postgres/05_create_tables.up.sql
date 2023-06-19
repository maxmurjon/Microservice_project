CREATE TABLE "product" (
  "id" VARCHAR PRIMARY KEY,
  "photo" VARCHAR(255) NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "category_id" VARCHAR NOT NULL REFERENCES "category"("id"),
  "branch_id" VARCHAR NOT NULL REFERENCES "branch"("id"),
  "barcode" VARCHAR(255) NOT NULL,
  "price" NUMERIC(10, 2) NOT NULL
);

-- ON DELETE CASCADE ON UPDATE NO ACTION