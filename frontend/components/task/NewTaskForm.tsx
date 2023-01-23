import { TaskForm } from '@/interfaces/task'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import dayjs from 'dayjs'
import Router from 'next/router'
import { useState } from 'react'
import Cookies from 'universal-cookie'


const NewTaskForm = () => {
  const [title, setTitle] = useState<string>("")
  const [description, setDescription] = useState<string>("")
  const [status, setStatus] = useState<string>("TODO")
  const [priority, setPriority] = useState<number>(1)
  const [dueDate, setDueDate] = useState<string>("")

  const handleSubmit = () => {
    const taskForm: TaskForm = {
      title: title,
      description: description,
      status: status,
      priority: priority,
      due_date: dayjs(dueDate).format("YYYY-MM-DDTHH:mm:ssZ"),
    }

    // get user_id from cookie
    const cookies = new Cookies()
    const userID = cookies.get('user_id')
    if (!userID) {
      Router.push('/login')
      return
    }

    if (title === '') {
      alert('Please type title')
      return
    }

    if (description === '') {
      alert('Please type description')
      return
    }

    if (dueDate === '') {
      alert('Please set deadline')
      return
    }

    axios.post(`${url}/v1/tasks/new`, taskForm, {
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((res: AxiosResponse<number>) => {
        const { data, status } = res
        switch (status) {
          case 200:
            alert("Task Created Successfully, Task ID: " + data)
            break
          case 400:
            alert("Bad Request" + res.statusText)
            break
          case 500:
            alert("Internal Server Error" + res.statusText)
            break
          default:
            alert("Something went wrong" + res.statusText)
        }
      }
      )
  }

  return (
    <div className="mt-10 sm:mt-0">
      <div className="md:grid md:grid-cols-3 md:gap-6">
        <div className="md:col-span-1">
          <div className="px-4 sm:px-0">
            <h3 className="text-lg font-medium leading-6 text-gray-900">Create a New Task</h3>
            <p className="mt-1 text-sm text-gray-600"> Please Type Your Task which you need to do. </p>
          </div>
        </div>
        <div className="mt-5 md:col-span-2 md:mt-0">
          <form action="#" method="POST">
            <div className="overflow-hidden shadow sm:rounded-md">
              <div className="bg-white px-4 py-5 sm:p-6">
                <div className="grid grid-cols-6 gap-6">
                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="user_name" className="block text-sm font-medium text-gray-700">Title</label>
                    <input type="text" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm tascholaForm" value={title} onChange={(event) => setTitle(event?.target.value)} />
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="description" className="block text-sm font-medium text-gray-700">Description</label>
                    <input type="text" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm h-20 tascholaForm" value={description} onChange={(event) => setDescription(event?.target.value)} />
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="status" className="block text-sm font-medium text-gray-700">Status</label>
                    <select className="mt-1 block w-full rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm tascholaForm" value={status} onChange={(event) => setStatus(event?.target.value)}>
                      <option value="TODO">TODO</option>
                      <option value="DOING">DOING</option>
                      <option value="DONE">DONE</option>
                    </select>
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="priority" className="block text-sm font-medium text-gray-700">Priority</label>
                    <select className="mt-1 block w-full rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm tascholaForm" value={priority} onChange={(event) => setPriority(Number(event?.target.value))}>
                      <option value="1">1 (Highest)</option>
                      <option value="2">2 (High)</option>
                      <option value="3">3 (Middle)</option>
                      <option value="4">4 (Low)</option>
                      <option value="5">5 (Lowest)</option>
                    </select>
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="due_date" className="block text-sm font-medium text-gray-700">Due Date</label>
                    <input type="date" className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" value={dueDate} onChange={(event) => setDueDate(event?.target.value)} />
                  </div>
                </div>
              </div>
              <div className="bg-gray-50 px-4 py-3 text-right sm:px-6">
                <button type="submit" className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2" onClick={handleSubmit}>Create</button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}

export default NewTaskForm
