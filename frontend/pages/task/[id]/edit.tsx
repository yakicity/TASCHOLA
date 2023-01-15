import EditTaskForm from '@/components/task/EditTaskForm'
import { Task } from '@/interfaces/task'
import styles from '@/styles/Home.module.scss'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'

const TaskEditPage = () => {
  const router = useRouter()
  const { id } = router.query

  const [task, setTask] = useState<Task>()
  const [loading, setLoading] = useState<boolean>(false)

  useEffect(() => {
    (async () => {
      await fetch(`http://localhost:8000/v1/task/${id}`)
        .then((res) => res.json())
        .then((res) => {
          setTask(res)
          setLoading(false)
        })
    })()
  }, [])

  return (
    <>
      <main className={styles.main}>
        {loading && <div>Loading...</div>}
        {!loading && !task && <div>No task found</div>}
        {!loading && task && <EditTaskForm task={task} />}
      </main>
    </>
  )
}

export default TaskEditPage
