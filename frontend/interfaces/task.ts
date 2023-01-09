export interface TaskForm {
  title: string;
  description: string;
  status: string
  priority: number
  due_date: string
}

export interface Task extends TaskForm {
  id: number;
  created_at: string;
}
