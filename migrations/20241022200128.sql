-- Create enum type "user_type"
CREATE TYPE "user_type" AS ENUM ('user', 'staff');
-- Create "course" table
CREATE TABLE "course" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_course_deleted_at" to table: "course"
CREATE INDEX "idx_course_deleted_at" ON "course" ("deleted_at");
-- Create "chapter" table
CREATE TABLE "chapter" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" character varying(80) NOT NULL,
  "order" integer NOT NULL,
  "course_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_course_chapters" FOREIGN KEY ("course_id") REFERENCES "course" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_chapter_deleted_at" to table: "chapter"
CREATE INDEX "idx_chapter_deleted_at" ON "chapter" ("deleted_at");
-- Create "section" table
CREATE TABLE "section" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" character varying(80) NOT NULL,
  "order" integer NOT NULL,
  "title" text NOT NULL,
  "content" text NOT NULL,
  "chapter_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_chapter_sections" FOREIGN KEY ("chapter_id") REFERENCES "chapter" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_section_deleted_at" to table: "section"
CREATE INDEX "idx_section_deleted_at" ON "section" ("deleted_at");
