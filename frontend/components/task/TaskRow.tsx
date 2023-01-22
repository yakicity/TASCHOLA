import { Task } from '@/interfaces/task'
import { url } from '@/utils/constants'
import axios from 'axios'
import Link from 'next/link'
import { getPriorityClassName, getStatusClassName, formatDateTime } from '@/functions/task'

type TaskRowProps = {
  task: Task
}

const TaskRow = (props: TaskRowProps) => {
  const { task } = props

  const handleSubmit = () => {
    try {
      axios.delete(`${url}/v1/tasks/${task.id}`, {
      }).then((res) => {
        alert('Task deleted successfully' + res.statusText)
      })
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <tr className="hover:bg-gray-50">
      <th className="flex gap-3 px-6 py-4 font-normal text-gray-900">
        <Link href={`/tasks/${task.id}`}>
          <div className="text-sm">
            <div className="font-medium text-gray-700">{task.title}
            </div>
            <div className="text-gray-400">{task.description}</div>
          </div>
        </Link>
      </th>
      <td className="px-6 py-4">
        <span
          className={`inline-flex items-center gap-1 rounded-full bg-${getStatusClassName(task.status)}-50 px-2 py-1 text-xs font-semibold text-${getStatusClassName(task.status)}-600`}
        >
          <span className={`h-1.5 w-1.5 rounded-full bg-${getStatusClassName(task.status)}-600`}></span>
          {task.status}
        </span>
      </td>
      <td className="px-6 py-4">{formatDateTime(task.due_date)}</td>
      <td className="px-6 py-4">
        <div className="gap-2 grid w-[10rem] grid-cols-5 rounded-xl bg-gray-200 p-1">
          <span
            className={`inline-flex text-center items-center justify-center gap-1 rounded-full ${getPriorityClassName(1, task.priority)} px-2 py-1 text-xs font-semibold`}
          >
            1
          </span>
          <span
            className={`inline-flex text-center items-center justify-center gap-1 rounded-full ${getPriorityClassName(2, task.priority)} px-2 py-1 text-xs font-semibold`}
          >
            2
          </span>
          <span
            className={`inline-flex text-center items-center justify-center gap-1 rounded-full ${getPriorityClassName(3, task.priority)} px-2 py-1 text-xs font-semibold`}
          >
            3
          </span>
          <span
            className={`inline-flex text-center items-center justify-center gap-1 rounded-full ${getPriorityClassName(4, task.priority)} px-2 py-1 text-xs font-semibold`}
          >
            4
          </span>
          <span
            className={`inline-flex text-center items-center justify-center gap-1 rounded-full ${getPriorityClassName(5, task.priority)} px-2 py-1 text-xs font-semibold`}
          >
            5
          </span>
        </div>
      </td>
      <td className="px-6 py-4">
        <div className="flex justify-end gap-4">
          <button onClick={handleSubmit}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth="1.5"
              stroke="currentColor"
              className="h-6 w-6"
              x-tooltip="tooltip"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0"
              />
            </svg>
          </button>
          <Link href={"/tasks/" + task.id + "/edit"}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth="1.5"
              stroke="currentColor"
              className="h-6 w-6"
              x-tooltip="tooltip"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L6.832 19.82a4.5 4.5 0 01-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 011.13-1.897L16.863 4.487zm0 0L19.5 7.125"
              />
            </svg>
          </Link>
        </div>
      </td>
    </tr>
  )
}

export default TaskRow
