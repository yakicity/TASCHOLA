-- +migrate Up
CREATE TABLE users (
  id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL UNIQUE,
  password BINARY(32) NOT NULL,
  PRIMARY KEY (id)
) DEFAULT CHARSET = utf8mb4;

-- +migrate Down
DROP TABLE users;
