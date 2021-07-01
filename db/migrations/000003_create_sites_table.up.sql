CREATE TABLE IF NOT EXISTS sites(
  id SERIAL PRIMARY KEY,
  domain VARCHAR (100) UNIQUE NOT NULL,
  name VARCHAR (100) NOT NULL,
  icon VARCHAR (100),
  publisher VARCHAR (100)
);

CREATE INDEX ON sites (domain);
CREATE INDEX ON sites (name);

-- TODO: 外部キー追加。データなし状態ではできないので後追いで行う
ALTER TABLE entries ADD COLUMN IF NOT EXISTS site_id INT NOT NULL DEFAULT 0;
