export interface StudentForm {
  user_id: number
  student_id: string
  password: string

  matrix_code_a_column: string
  matrix_code_b_column: string
  matrix_code_c_column: string
  matrix_code_d_column: string
  matrix_code_e_column: string
  matrix_code_f_column: string
  matrix_code_g_column: string
  matrix_code_h_column: string
  matrix_code_i_column: string
  matrix_code_j_column: string
}

export interface Student extends StudentForm {
  id: number
}
