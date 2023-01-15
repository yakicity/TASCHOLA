import { Task } from '@/interfaces/task'
import Link from 'next/link'

type TaskProps = {
  task: Task
}
const DetailTask = (props: TaskProps) => {
  const { task } = props

  return (
    <div className="bg-white">
      <div className="pt-6">
        <nav aria-label="Breadcrumb">
          <ol role="list" className="mx-auto flex max-w-2xl items-center space-x-2 px-4 sm:px-6 lg:max-w-7xl lg:px-8">
            <li>
              <div className="flex items-center">
                <a href="/tasks" className="mr-2 text-sm font-medium text-gray-900">task</a>
                <svg width="16" height="20" viewBox="0 0 16 20" fill="currentColor" xmlns="http://www.w3.org/2000/svg" aria-hidden="true" className="h-5 w-4 text-gray-300">
                  <path d="M5.697 4.34L8.98 16.532h1.327L7.025 4.341H5.697z" />
                </svg>
              </div>
            </li>

            <li className="text-sm">
              <a href="#" aria-current="page" className="font-medium text-gray-500 hover:text-gray-600">{task.id}</a>
            </li>
          </ol>
        </nav>

        <div className="mx-auto max-w-2xl px-4 pt-10 pb-16 sm:px-6 lg:grid lg:max-w-7xl lg:grid-cols-3 lg:grid-rows-[auto,auto,1fr] lg:gap-x-8 lg:px-8 lg:pt-16 lg:pb-24">
          <div className="lg:col-span-2 lg:border-r lg:border-gray-200 lg:pr-8">
            <h1 className="text-2xl font-bold tracking-tight text-gray-900 sm:text-3xl">{task.title}</h1>
          </div>

          <div className="mt-4 lg:row-span-3 lg:mt-0">
            <h2 className="sr-only">Due Date</h2>
            <h3 className="text-sm font-medium text-gray-900">Due Date</h3>
            <p className="text-3xl tracking-tight text-gray-900">{task.due_date}</p>

            <form className="mt-10">
              <div>
                <h3 className="text-sm font-medium text-gray-900">Status</h3>
                <fieldset className="mt-4">
                  <legend className="sr-only">Status</legend>
                  <div className="flex items-center space-x-3">
                    <label className="-m-0.5 relative p-0.5 rounded-full flex items-center justify-center cursor-pointer focus:outline-none ring-gray-400">
                      <input type="radio" name="color-choice" value="White" className="sr-only" aria-labelledby="color-choice-0-label" />
                      <span id="color-choice-0-label" className="sr-only"> White </span>
                      <span aria-hidden="true" className="h-8 w-8 bg-white border border-black border-opacity-10 rounded-full"></span>
                    </label>

                    <label className="-m-0.5 relative p-0.5 rounded-full flex items-center justify-center cursor-pointer focus:outline-none ring-gray-400">
                      <input type="radio" name="color-choice" value="Gray" className="sr-only" aria-labelledby="color-choice-1-label" />
                      <span id="color-choice-1-label" className="sr-only"> Gray </span>
                      <span aria-hidden="true" className="h-8 w-8 bg-gray-200 border border-black border-opacity-10 rounded-full"></span>
                    </label>

                    <label className="-m-0.5 relative p-0.5 rounded-full flex items-center justify-center cursor-pointer focus:outline-none ring-gray-900">
                      <input type="radio" name="color-choice" value="Black" className="sr-only" aria-labelledby="color-choice-2-label" />
                      <span id="color-choice-2-label" className="sr-only"> Black </span>
                      <span aria-hidden="true" className="h-8 w-8 bg-gray-900 border border-black border-opacity-10 rounded-full"></span>
                    </label>
                  </div>
                </fieldset>
              </div>

              <div>
                <h3 className="text-sm font-medium text-gray-900">Priority</h3>
                <fieldset className="mt-4">
                  <legend className="sr-only">Status</legend>
                  <div className="flex items-center space-x-3">
                    <label className="-m-0.5 relative p-0.5 rounded-full flex items-center justify-center cursor-pointer focus:outline-none ring-gray-400">
                      <input type="radio" name="color-choice" value="White" className="sr-only" aria-labelledby="color-choice-0-label" />
                      <span id="color-choice-0-label" className="sr-only"> White </span>
                      <span aria-hidden="true" className="h-8 w-8 bg-white border border-black border-opacity-10 rounded-full"></span>
                    </label>

                    <label className="-m-0.5 relative p-0.5 rounded-full flex items-center justify-center cursor-pointer focus:outline-none ring-gray-400">
                      <input type="radio" name="color-choice" value="Gray" className="sr-only" aria-labelledby="color-choice-1-label" />
                      <span id="color-choice-1-label" className="sr-only"> Gray </span>
                      <span aria-hidden="true" className="h-8 w-8 bg-gray-200 border border-black border-opacity-10 rounded-full"></span>
                    </label>

                    <label className="-m-0.5 relative p-0.5 rounded-full flex items-center justify-center cursor-pointer focus:outline-none ring-gray-900">
                      <input type="radio" name="color-choice" value="Black" className="sr-only" aria-labelledby="color-choice-2-label" />
                      <span id="color-choice-2-label" className="sr-only"> Black </span>
                      <span aria-hidden="true" className="h-8 w-8 bg-gray-900 border border-black border-opacity-10 rounded-full"></span>
                    </label>
                  </div>
                </fieldset>
              </div>
              <Link href={`/task/${task.id}/edit`}>
                <button type="submit" className="mt-10 flex w-full items-center justify-center rounded-md border border-transparent bg-indigo-600 py-3 px-8 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">Edit</button>
              </Link>

              <button type="submit" className="mt-10 flex w-full items-center justify-center rounded-md border border-transparent bg-indigo-600 py-3 px-8 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">Delete</button>
            </form>
          </div>

          <div className="py-10 lg:col-span-2 lg:col-start-1 lg:border-r lg:border-gray-200 lg:pt-6 lg:pb-16 lg:pr-8">
            <div>
              <h3 className="sr-only">Description</h3>

              <div className="space-y-6">
                <p className="text-base text-gray-900">{task.description}</p>
              </div>
            </div>

            <div className="mt-10">
              <h3 className="text-sm font-medium text-gray-900">Created At</h3>

              <div className="mt-4">
                <ul role="list" className="list-disc space-y-2 pl-4 text-sm">
                  <li className="text-gray-400"><span className="text-gray-600">{task.created_at}</span></li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default DetailTask
