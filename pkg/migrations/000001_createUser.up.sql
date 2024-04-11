CREATE TABLE users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    login text NOT NULL,
    name text NULL,
    phone text NOT NULL
);  
CREATE INDEX IF NOT EXISTS "idx_users_id" ON "users" ("id");
