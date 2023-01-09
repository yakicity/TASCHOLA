import TaskList from '@/components/task/TaskList'
import { Task } from '@/interfaces/task'
import styles from '@/styles/Home.module.scss'

const taskList = [
  {
    id: 1,
    title: 'Task 1',
    description: 'Task 1 description test test test test .....................',
    status: 'TODO',
    created_at: '2020-01-01',
    priority: 1,
    due_date: '2021-01-01',
  } as Task,
  {
    id: 2,
    title: 'Task 2',
    description: 'Task 2 description',
    status: 'TODO',
    created_at: '2020-01-01',
    priority: 1,
    due_date: '2021-01-01',
  } as Task,
]

const TasksPage = () => {
  return (
    <>
      <main className={styles.main}>
        <h1>Task List</h1>
        <div className="flex justify-end gap-4">
          <a href="/task/new" className="btn btn-primary">
            Create Task
          </a>
        </div>
        <TaskList tasks={taskList} />
      </main>
    </>
  )
}

export default TasksPage
