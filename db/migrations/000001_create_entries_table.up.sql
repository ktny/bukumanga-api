CREATE FUNCTION set_update_time() RETURNS OPAQUE AS '
  begin
    new.updated_at := ''now'';
    return new;
  end;
' LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS entries(
  id BIGINT PRIMARY KEY,
  title VARCHAR (300) UNIQUE NOT NULL,
  url VARCHAR (300) UNIQUE NOT NULL,
  domain VARCHAR (100) NOT NULL,
  bookmark_count INT NOT NULL,
  image VARCHAR (1000),
  published_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_tri BEFORE UPDATE ON entries FOR EACH ROW EXECUTE PROCEDURE set_update_time();
