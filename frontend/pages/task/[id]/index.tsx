import DetailTask from '@/components/task/DetailTask'
import { Task } from '@/interfaces/task'
import styles from '@/styles/Home.module.scss'
import axios from 'axios'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'

const TaskPage = () => {
  const router = useRouter()
  const { id } = router.query

  const [task, setTask] = useState<Task>()
  const [loading, setLoading] = useState<boolean>(false)

  useEffect(() => {
    (async () => {
      await axios.get(`http://localhost:8000/v1/task/${id}`)
        .then((res) => {
          setTask(res.data)
          setLoading(false)
        }
        )
    })()
  }, [])

  return (
    <>
      <main className={styles.main}>
        {loading && <div>Loading...</div>}
        {!loading && !task && <div>No task found</div>}
        {!loading && task && <DetailTask task={task} />}
      </main>
    </>
  )
}

export default TaskPage
