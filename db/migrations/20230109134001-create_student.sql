-- +migrate Up
CREATE TABLE students (
  id BIGINT NOT NULL AUTO_INCREMENT,
  user_id BIGINT NOT NULL,
  student_id VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  matrix_code_a_column VARCHAR(255) NOT NULL,
  matrix_code_b_column VARCHAR(255) NOT NULL,
  matrix_code_c_column VARCHAR(255) NOT NULL,
  matrix_code_d_column VARCHAR(255) NOT NULL,
  matrix_code_e_column VARCHAR(255) NOT NULL,
  matrix_code_f_column VARCHAR(255) NOT NULL,
  matrix_code_g_column VARCHAR(255) NOT NULL,
  matrix_code_h_column VARCHAR(255) NOT NULL,
  matrix_code_i_column VARCHAR(255) NOT NULL,
  matrix_code_j_column VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id)
) DEFAULT CHARSET = utf8mb4;

-- +migrate Down
DROP TABLE students;
