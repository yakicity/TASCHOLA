import { AuthRequest, AuthResponse } from '@/interfaces/auth'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import { useState } from "react"
import Cookies from 'universal-cookie'

const LoginForm = () => {
  const [userName, setUserName] = useState<string>('')
  const [password, setPassword] = useState<string>('')

  const handleSubmit = () => {
    const data: AuthRequest = {
      name: userName,
      password: password,
    }

    if (userName === '') {
      alert('Please type user name')
      return
    }

    try {
      console.log("login post:" + data.name + " " + data.password)
      axios.post(`${url}/v1/login`, data, {
        headers: {
          'Content-Type': 'application/json',
        }
      })
        .then((res: AxiosResponse<AuthResponse>) => {
          const { data, status } = res
          switch (status) {
            case 200:
              const cookies = new Cookies()
              cookies.set('user_id', data.user_id, { path: '/' })
              window.location.href = '/'
              break
            default:
              alert('Something went wrong' + res.statusText)
          }
        })
    } catch (error) {
      alert(console.error())
    }
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
                    <input type="text" name="user_name" id="user_name" autoComplete="user_name" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm tascholaForm" value={userName} onChange={(event) => setUserName(event?.target.value)} />
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="password1" className="block text-sm font-medium text-gray-700">Password</label>
                    <input type="password" name="password1" id="password1" autoComplete="password" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm tascholaForm" value={password} onChange={(event) => setPassword(event?.target.value)} />
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
