CREATE TABLE IF NOT EXISTS publishers(
  id SERIAL PRIMARY KEY,
  domain VARCHAR (100) UNIQUE NOT NULL,
  name VARCHAR (100) NOT NULL,
  icon VARCHAR (100),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON publishers (name);

-- TODO: 外部キー追加。データなし状態ではできないので後追いで行う
ALTER TABLE entries ADD COLUMN IF NOT EXISTS publisher_id INT NOT NULL DEFAULT 0;

CREATE INDEX ON entries (publisher_id);
