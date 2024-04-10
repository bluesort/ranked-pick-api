BEGIN;

DROP TYPE IF EXISTS survey_state;

ALTER TABLE surveys
  RENAME COLUMN description TO prompt;
  DROP COLUMN title;
  DROP COLUMN state;

COMMIT;
