ALTER TABLE user
    DROP COLUMN lastname;

ALTER TABLE user ADD COLUMN blub;

ALTER TABLE venue ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW();