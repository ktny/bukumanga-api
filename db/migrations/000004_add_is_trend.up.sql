ALTER TABLE entries ADD COLUMN IF NOT EXISTS is_trend BOOLEAN NOT NULL DEFAULT false;
CREATE INDEX ON entries (is_trend);
