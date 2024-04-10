BEGIN;

CREATE TYPE survey_state AS ENUM (
  'pending',
  'gathering_options',
  'voting',
  'closed',
  'deleted'
);

ALTER TABLE surveys RENAME COLUMN prompt TO description;
ALTER TABLE surveys ADD COLUMN title VARCHAR(300) NOT NULL;
ALTER TABLE surveys ADD COLUMN state survey_state NOT NULL DEFAULT 'pending';

COMMIT;
