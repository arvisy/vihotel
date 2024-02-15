CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL UNIQUE,
  "password" VARCHAR NOT NULL,
  "role_id" INT NOT NULL,
  "address_id" INT,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL
);

CREATE TABLE "address" (
  "id" SERIAL PRIMARY KEY,
  "country" VARCHAR NOT NULL,
  "city" VARCHAR NOT NULL
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "users" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");