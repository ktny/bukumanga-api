CREATE TABLE IF NOT EXISTS comments(
  id SERIAL PRIMARY KEY,
  entry_id INT NOT NULL REFERENCES entries (id),
  rank SMALLINT NOT NULL CHECK (rank > 0),
  username VARCHAR (100) NOT NULL,
  icon VARCHAR (100) NOT NULL,
  content VARCHAR (200) NOT NULL,
  commented_at DATE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (entry_id, rank),
  UNIQUE (entry_id, username)
);
