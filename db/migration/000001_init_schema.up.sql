
CREATE TABLE "reviews" (
  "id" bigserial PRIMARY KEY,
  "user_id" int,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "date" timestamptz NOT NULL,
  "classification" varchar NOT NULL,
  "amount" numeric,
  "overall_rate" int NOT NULL,
  "insurance" varchar,
  "place_id" int NOT NULL,
  "place_rate" int NOT NULL,
  "doctor_id" int NOT NULL,
  "doctor_rate" int NOT NULL,
  "midwife_id" int,
  "midwife_rate" int,
  "doula_id" int,
  "doula_rate" int,
  "team" varchar,
  "team_rate" int,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz,
  "image" varchar
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "doctors" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "crm" varchar NOT NULL,
  "average_rate" int NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "places" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "address" varchar,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "country" varchar,
  "average_rate" int NOT NULL,
  "classification" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "midwifes" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "average_rate" int NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "doula" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "average_rate" int NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "reviews" ("doctor_id");

CREATE INDEX ON "reviews" ("place_id");

CREATE INDEX ON "reviews" ("overall_rate");

CREATE INDEX ON "doctors" ("name");

CREATE INDEX ON "doctors" ("average_rate");

CREATE INDEX ON "places" ("name");

CREATE INDEX ON "places" ("average_rate");

CREATE INDEX ON "places" ("city");

CREATE INDEX ON "midwifes" ("name");

CREATE INDEX ON "midwifes" ("average_rate");

CREATE INDEX ON "doula" ("name");

CREATE INDEX ON "doula" ("average_rate");

ALTER TABLE "reviews" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("place_id") REFERENCES "places" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("doctor_id") REFERENCES "doctors" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("midwife_id") REFERENCES "midwifes" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("doula_id") REFERENCES "doula" ("id");
