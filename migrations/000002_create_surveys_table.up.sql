CREATE TABLE IF NOT EXISTS surveys(
  id INTEGER PRIMARY KEY NOT NULL,
  user_id INTEGER NOT NULL,
  title VARCHAR(300) NOT NULL,
  state TEXT CHECK( state IN ('pending', 'gathering_options', 'voting', 'closed', 'deleted') ) NOT NULL DEFAULT 'pending',
  visibility TEXT CHECK( visibility IN ('public', 'link', 'private') ) NOT NULL DEFAULT 'public',

  description VARCHAR(300),

  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE UNIQUE INDEX idx_surveys_state_user ON surveys (state, user_id);
