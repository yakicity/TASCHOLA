-- +migrate Up
CREATE TABLE ownerships (
  id BIGINT NOT NULL AUTO_INCREMENT,
  user_id BIGINT NOT NULL,
  task_id BIGINT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (task_id) REFERENCES tasks(id)
) DEFAULT CHARSET = utf8mb4;

-- +migrate Down
DROP TABLE ownerships;
