import axios from "axios"
import { useState } from "react"

const LoginForm = () => {
  const [userName, setUserName] = useState<string>('')
  const [password, setPassword] = useState<string>('')

  const handleSubmit = () => {
    const data = {
      user_name: userName,
      password: password,
    }

    if (userName === '') {
      alert('Please type user name')
      return
    }

    axios.post('http://localhost:8000/v1/login', data)
      .then((res) => {
        console.log(res)
      })
      .catch((err) => {
        console.log(err)
      })
  }

  return (
    <div className="mt-10 sm:mt-0">
      <div className="md:grid md:grid-cols-3 md:gap-6">
        <div className="md:col-span-1">
          <div className="px-4 sm:px-0">
            <h3 className="text-lg font-medium leading-6 text-gray-900">Login</h3>
            <p className="mt-1 text-sm text-gray-600"> Please type your user name and password. </p>
          </div>
        </div>
        <div className="mt-5 md:col-span-2 md:mt-0">
          <form action="#" method="POST">
            <div className="overflow-hidden shadow sm:rounded-md">
              <div className="bg-white px-4 py-5 sm:p-6">
                <div className="grid grid-cols-6 gap-6">
                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="user_name" className="block text-sm font-medium text-gray-700">User Name</label>
                    <input type="text" name="user_name" id="user_name" autoComplete="user_name" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" value={userName} onChange={(event) => setUserName(event?.target.value)} />
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="password1" className="block text-sm font-medium text-gray-700">Password</label>
                    <input type="text" name="password1" id="password1" autoComplete="password" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" value={password} onChange={(event) => setPassword(event?.target.value)} />
                  </div>
                </div>
              </div>
              <div className="bg-gray-50 px-4 py-3 text-right sm:px-6">
                <button type="submit" className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2" onClick={handleSubmit}>Login</button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}

export default LoginForm
