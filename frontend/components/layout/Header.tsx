import Link from 'next/link'

const Header = () => {
  return (
    <header>
      <div className="relative bg-white">
        <div className="mx-auto max-w-7xl px-6">
          <div className="flex items-center justify-between border-b-2 border-gray-100 py-6 md:justify-start md:space-x-10">
            <div className="flex justify-start lg:w-0 lg:flex-1">
              <Link href="/">
                <span className="sr-only">Your Company</span>
                <img className="h-8 w-auto sm:h-10" src="/taschola-logo.svg" alt="" />
              </Link>
            </div>

            <nav className="hidden space-x-10 md:flex">
              <Link href="/tasks" className="text-base font-medium text-gray-500 hover:text-gray-900">Tasks</Link>
              <Link href="/user" className="text-base font-medium text-gray-500 hover:text-gray-900">Profile</Link>
              <Link href="" className="text-base font-medium text-gray-500 hover:text-gray-900">Sync</Link>
            </nav>

            <div className="hidden items-center justify-end md:flex md:flex-1 lg:w-0">
              <Link href="/login" className="whitespace-nowrap text-base font-medium text-gray-500 hover:text-gray-900">Login</Link>
              <Link href="/signup" className="ml-8 inline-flex items-center justify-center whitespace-nowrap rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-indigo-700">Sign up</Link>
            </div>
          </div>
        </div>
      </div>
    </header>
  )
}

export default Header
