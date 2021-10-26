CREATE TABLE IF NOT EXISTS invoice (
  id serial primary key,
  message text NOT NULL,
  created_at timestamp NOT NULL
);