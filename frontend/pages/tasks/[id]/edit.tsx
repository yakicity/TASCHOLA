import EditTaskForm from '@/components/task/EditTaskForm'
import { Task } from '@/interfaces/task'
import styles from '@/styles/Home.module.scss'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import Router, { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import Cookies from 'universal-cookie'

const TaskEditPage = () => {
  const router = useRouter()
  const { id } = router.query

  const [task, setTask] = useState<Task>()
  const [loading, setLoading] = useState<boolean>(false)

  useEffect(() => {
    if (!router.isReady) {
      return
    }

    // Check if user is logged in
    const cookies = new Cookies()
    const userID = cookies.get('user_id')
    if (!userID) {
      Router.push('/login')
      return
    }

    setLoading(true)

    // API Request
    axios.get(`${url}/v1/tasks/${id}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      }
    })
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
  }, [router.query])

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
