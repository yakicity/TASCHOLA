import TaskList from '@/components/task/TaskList'
import { Task } from '@/interfaces/task'
import styles from '@/styles/Home.module.scss'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import Link from 'next/link'
import Router from 'next/router'
import { useEffect, useState } from 'react'
import Cookies from 'universal-cookie'

const TasksPage = () => {
  const [tasks, setTasks] = useState<Task[]>([])
  const [loading, setLoading] = useState<boolean>(true)

  useEffect(() => {
    // Check if user is logged in
    const cookies = new Cookies()
    const userID = cookies.get('user_id')
    if (!userID) {
      Router.push('/login')
      return
    }

    // API Request
    try {
      axios.get(`${url}/v1/tasks/`)
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
              alert('Something went wrong')
          }
        })
    } catch (error) {
      // redirect to login page
      Router.push('/user/login')
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
          {!loading && tasks != null && tasks.length === 0 && <div>No tasks found</div>}
          {!loading && tasks != null && tasks.length > 0 && <TaskList tasks={tasks} />}
        </div>
      </main>
    </>
  )
}

export default TasksPage
