BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS survey_options (
  id INTEGER PRIMARY KEY NOT NULL,
	survey_id INTEGER NOT NULL,
	title VARCHAR(300) NOT NULL,

  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY (survey_id) REFERENCES surveys(id)
);

CREATE TABLE IF NOT EXISTS survey_answers (
  id INTEGER PRIMARY KEY NOT NULL,
	survey_id INTEGER NOT NULL,
	survey_option_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	rank INTEGER NOT NULL,

  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (survey_id) REFERENCES surveys(id),
	FOREIGN KEY (survey_option_id) REFERENCES survey_options(id)
);
CREATE UNIQUE INDEX idx_survey_answers_on_user_option ON survey_answers (user_id, survey_option_id);

COMMIT;
