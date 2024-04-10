CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  display_name VARCHAR (50),
  password VARCHAR (50) NOT NULL,
  email VARCHAR (300) UNIQUE NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);
