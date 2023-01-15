import { Task } from '@/interfaces/task'

type EditTaskFormProps = {
  task: Task
}

const EditTaskForm = (props: EditTaskFormProps) => {
  const { task } = props

  return (
    <div className="mt-10 sm:mt-0">
      <div className="md:grid md:grid-cols-3 md:gap-6">
        <div className="md:col-span-1">
          <div className="px-4 sm:px-0">
            <h3 className="text-lg font-medium leading-6 text-gray-900">Edit Your task Task</h3>
          </div>
        </div>
        <div className="mt-5 md:col-span-2 md:mt-0">
          <form action="#" method="POST">
            <div className="overflow-hidden shadow sm:rounded-md">
              <div className="bg-white px-4 py-5 sm:p-6">
                <div className="grid grid-cols-6 gap-6">
                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="user_name" className="block text-sm font-medium text-gray-700">Title</label>
                    {/* old value */}
                    <p className="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">{task.title}</p>
                    <input type="text" name="user_name" id="user_name" autoComplete="user_name" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" />
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="password1" className="block text-sm font-medium text-gray-700">Description</label>
                    {/* old value */}
                    <p className="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">{task.description}</p>
                    <input type="text" name="password1" id="password1" autoComplete="password" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm h-20" />
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="password2" className="block text-sm font-medium text-gray-700">Status</label>
                    {/* old value */}
                    <p className="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">{task.status}</p>
                    <select id="student_type" name="student_type" autoComplete="degree" className="mt-1 block w-full rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm">
                      <option>TODO</option>
                      <option>DOING</option>
                      <option>DONE</option>
                    </select>
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="student_type" className="block text-sm font-medium text-gray-700">Priority</label>
                    {/* old value */}
                    <p className="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">{task.priority}</p>
                    <select id="student_type" name="student_type" autoComplete="degree" className="mt-1 block w-full rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm">
                      <option>1 (Highest)</option>
                      <option>2 (High)</option>
                      <option>3 (Middle)</option>
                      <option>4 (Low)</option>
                      <option>5 (Lowest)</option>
                    </select>
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="student_type" className="block text-sm font-medium text-gray-700">Due Date</label>
                    {/* old value */}
                    <p className="mt-1 block w-full rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">{task.due_date}</p>
                    <input type="date" name="student_type" id="student_type" autoComplete="degree" className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" />
                  </div>
                </div>
              </div>
              <div className="bg-gray-50 px-4 py-3 text-right sm:px-6">
                <button type="submit" className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">Update</button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}

export default EditTaskForm
