import TaskRow from '@/components/task/TaskRow'
import { Task } from '@/interfaces/task'

type TaskListProps = {
  tasks: Task[]
}

const TaskList = (props: TaskListProps) => {
  const { tasks } = props
  return (
    <div className="overflow-hidden rounded-lg border border-gray-200 shadow-md m-5">
      <table className="w-full border-collapse bg-white text-left text-sm text-gray-500">
        <thead className="bg-gray-50">
          <tr>
            <th scope="col" className="px-6 py-4 font-medium text-gray-900">Title</th>
            <th scope="col" className="px-6 py-4 font-medium text-gray-900">Status</th>
            <th scope="col" className="px-6 py-4 font-medium text-gray-900">Due Date</th>
            <th scope="col" className="px-6 py-4 font-medium text-gray-900">Priority</th>
            <th scope="col" className="px-6 py-4 font-medium text-gray-900"></th>
          </tr>
        </thead>
        <tbody className="divide-y divide-gray-100 border-t border-gray-100">
          {tasks.map((task) => {
            return <TaskRow task={task} key={task.id} />
          }
          )}
        </tbody>
      </table>
    </div>
  )
}

export default TaskList
