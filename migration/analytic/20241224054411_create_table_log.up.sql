BEGIN;


CREATE TABLE IF NOT EXISTS paste_log (
  time TIMESTAMPTZ NOT NULL,
  shortlink text not null
);


SELECT create_hypertable('paste_log', by_range('time'));

COMMIT;