
-- +migrate Up
CREATE TABLE authors (id uuid PRIMARY KEY);

-- +migrate Down
DROP TABLE authors;
