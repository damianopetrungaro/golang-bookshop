
-- +migrate Up
CREATE TABLE books (
  id UUID PRIMARY KEY,
  title VARCHAR(50),
  year DECIMAL(4),
  isbn VARCHAR(13),
  author_id UUID,
  created_at DATE,
  updated_at DATE NULLABLE
);

-- +migrate Down
DROP TABLE books;
