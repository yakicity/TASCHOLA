import TaskList from '@/components/task/TaskList'
import { Task } from '@/interfaces/task'
import styles from '@/styles/Home.module.scss'
import Link from 'next/link'

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
        <div>
          <div>
            <Link href="/task/new" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">
              Create a new Task
            </Link>
          </div>

          <div>
            <TaskList tasks={taskList} />
          </div>
        </div>
      </main>
    </>
  )
}

export default TasksPage
