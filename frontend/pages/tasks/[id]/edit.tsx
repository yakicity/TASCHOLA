import EditTaskForm from '@/components/task/EditTaskForm'
import { Task } from '@/interfaces/task'
import styles from '@/styles/Home.module.scss'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'


const TaskEditPage = () => {
  const router = useRouter()
  const { id } = router.query

  const [task, setTask] = useState<Task>()
  const [loading, setLoading] = useState<boolean>(false)

  useEffect(() => {
    (async () => {
      setLoading(true)
      await axios.get(`${url}/v1/tasks/${id}`)
        .then((res: AxiosResponse<Task>) => {
          const { data, status } = res
          switch (status) {
            case 200:
              setTask(data)
              setLoading(false)
              break
            case 404:
              alert('Task not found' + res.statusText)
              break
            default:
              alert('Something went wrong' + res.statusText)
          }
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
