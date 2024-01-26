CREATE TABLE IF NOT EXISTS fijoy_workspace (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  namespace TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT now (),
  CONSTRAINT "fijoy_workspace_namespace_unique" UNIQUE ("namespace")
);

--> statement-breakpoint
CREATE INDEX IF NOT EXISTS "fijoy_workspace_namespace_index" ON "fijoy_workspace" ("namespace");