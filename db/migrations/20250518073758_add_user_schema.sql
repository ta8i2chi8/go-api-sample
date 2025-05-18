-- Create "users" table
CREATE TABLE "public"."users" ("id" serial NOT NULL, "name" text NOT NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
