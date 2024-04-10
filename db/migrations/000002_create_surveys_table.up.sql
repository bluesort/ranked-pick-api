CREATE TABLE IF NOT EXISTS surveys(
  user_id serial PRIMARY KEY,
  prompt VARCHAR (300),
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);
