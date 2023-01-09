export interface UserForm {
  name: string;
  password: string;
}

export interface User {
  id: number;
  name: string;
  password: number // byte[]
}
