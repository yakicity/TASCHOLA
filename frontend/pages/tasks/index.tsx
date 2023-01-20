import TaskList from '@/components/task/TaskList'
import { Task } from '@/interfaces/task'
import styles from '@/styles/Home.module.scss'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import Link from 'next/link'
import { useEffect, useState } from 'react'

const TasksPage = () => {
  const [tasks, setTasks] = useState<Task[]>([])
  const [loading, setLoading] = useState<boolean>(false)

  useEffect(() => {
    try {
      setLoading(true)
      axios.get(`${url}/v1/tasks/`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        }
      },)
        .then((res: AxiosResponse<Task[]>) => {
          const { data, status } = res
          switch (status) {
            case 200:
              setTasks(data)
              setLoading(false)
              break
            case 404:
              alert('Tasks not found' + res.statusText)
              break
            default:
              alert('Something went wrong' + res.statusText)
          }
        })
    } catch (error) {
      console.log(error)
    }
  }, [])

  return (
    <>
      <main className={styles.main}>
        <div>
          <div>
            <Link href="/tasks/new" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">
              Create a new Task
            </Link>
          </div>

          {loading && <div>Loading...</div>}
          {!loading && tasks.length === 0 && <div>No tasks found</div>}
          {!loading && tasks.length > 0 && <TaskList tasks={tasks} />}
        </div>
      </main>
    </>
  )
}

export default TasksPage
