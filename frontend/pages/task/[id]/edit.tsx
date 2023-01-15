import EditTaskForm from '@/components/task/EditTaskForm'
import { Task } from '@/interfaces/task'
import styles from '@/styles/Home.module.scss'
import { useRouter } from 'next/router'

const task = {
  id: 1,
  title: 'Task 1',
  description: 'Task 1 description test test test test .....................test test sentence',
  status: 'TODO',
  created_at: '2020-01-01',
  priority: 1,
  due_date: '2021-01-01',
} as Task

const TaskEditPage = () => {
  const router = useRouter()
  const { id } = router.query

  return (
    <>
      <main className={styles.main}>
        <EditTaskForm task={task} />
      </main>
    </>
  )
}

export default TaskEditPage
