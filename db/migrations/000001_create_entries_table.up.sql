CREATE TABLE IF NOT EXISTS entries(
  id SERIAL PRIMARY KEY,
  title VARCHAR (300) NOT NULL,
  url VARCHAR (300) UNIQUE NOT NULL,
  domain VARCHAR (100) NOT NULL,
  bookmark_count SMALLINT NOT NULL CHECK (bookmark_count >= 0),
  image VARCHAR (1000),
  hotentried_at DATE NOT NULL,
  published_at DATE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON entries (title);
CREATE INDEX ON entries (domain);
CREATE INDEX ON entries (bookmark_count);
CREATE INDEX ON entries (hotentried_at);
CREATE INDEX ON entries (published_at);

CREATE FUNCTION set_update_time() RETURNS TRIGGER AS '
  begin
    new.updated_at := ''now'';
    return new;
  end;
' LANGUAGE plpgsql;

CREATE TRIGGER update_tri BEFORE UPDATE ON entries FOR EACH ROW EXECUTE PROCEDURE set_update_time();
