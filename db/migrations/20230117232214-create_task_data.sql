
-- +migrate Up
INSERT INTO
  tasks (id, title, description, status, priority, created_at, due_date)
VALUES
  (1, 'Task 1', 'Description 1', 'TODO', 1, '2021-01-01', '2021-01-24'),
  (2, 'Task 2', 'Description 2', 'DONE', 3, '2021-01-01', '2021-01-25'),
  (3, 'Task 3', 'Description 3', 'DOING', 5, '2021-01-01', '2021-02-01'),
  (4, 'Task 4', 'Description 4', 'TODO', 4, '2021-01-01', '2021-03-01'),
  (5, 'Task 5', 'Description 5', 'TODO', 3, '2021-01-01', '2021-03-01');

INSERT INTO
  ownerships (user_id, task_id)
VALUES
  (1, 1),
  (1, 2),
  (1, 3),
  (2, 4),
  (2, 5);
-- +migrate Down
DELETE FROM ownerships WHERE task_id IN (1, 2, 3, 4, 5);

DELETE FROM tasks WHERE id IN (1, 2, 3, 4, 5);
