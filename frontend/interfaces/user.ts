export interface UserForm {
  name: string;
  password: string;
}

export interface User {
  id: number;
  name: string;
  password: number // byte[]
}

export interface UpdateUserForm {
  old_password: string;
  user: UserForm;
}
