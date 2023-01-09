const Header = () => {
  return (
    <header>
      <div className="relative bg-white">
        <div className="mx-auto max-w-7xl px-6">
          <div className="flex items-center justify-between border-b-2 border-gray-100 py-6 md:justify-start md:space-x-10">
            <div className="flex justify-start lg:w-0 lg:flex-1">
              <a href="/">
                <span className="sr-only">Your Company</span>
                <img className="h-8 w-auto sm:h-10" src="/taschola.svg" alt="" />
              </a>
            </div>

            <nav className="hidden space-x-10 md:flex">
              <a href="/tasks" className="text-base font-medium text-gray-500 hover:text-gray-900">Tasks</a>
              <a href="/user" className="text-base font-medium text-gray-500 hover:text-gray-900">Profile</a>
              <a href="" className="text-base font-medium text-gray-500 hover:text-gray-900">Sync</a>
            </nav>

            <div className="hidden items-center justify-end md:flex md:flex-1 lg:w-0">
              <a href="/login" className="whitespace-nowrap text-base font-medium text-gray-500 hover:text-gray-900">Login</a>
              <a href="/signup" className="ml-8 inline-flex items-center justify-center whitespace-nowrap rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-indigo-700">Sign up</a>
            </div>
          </div>
        </div>
      </div>
    </header>
  )
}

export default Header
