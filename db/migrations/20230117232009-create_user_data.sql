-- +migrate Up
INSERT INTO
  users (id, name, password)
VALUES
  (1, 'admin', 'admin-password'),
  (2, 'user', 'user-password');

-- +migrate Down
DELETE FROM users WHERE id IN (1, 2);
