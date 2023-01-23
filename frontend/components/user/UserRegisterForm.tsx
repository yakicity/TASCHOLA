import { AuthRequest, AuthResponse } from '@/interfaces/auth'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import { useState } from 'react'

const UserRegisterForm = () => {
  const [userName, setUserName] = useState<string>('')
  const [password1, setPassword1] = useState<string>('')
  const [password2, setPassword2] = useState<string>('')

  const handleSubmit = () => {
    const data: AuthRequest = {
      name: userName,
      password: password1
    }

    if (password1 !== password2) {
      console.log(password1, password2) // for debug
      alert('Password is not match')
      return
    }

    if (userName === '') {
      alert('Please type user name')
      return
    }

    axios.post(`${url}/v1/user/new`, data, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
      .then((res: AxiosResponse<AuthResponse>) => {
        const { data, status } = res
        console.log(data, status) // for debug
        switch (status) {
          case 200:
            if (!confirm('User created successfully')) {
              // redirect to login page
              window.location.href = '/'
            }

            break
          default:
            alert('Something went wrong' + res.statusText)
        }
      })
  }

  return (
    <div className="mt-10 sm:mt-0">
      <div className="md:grid md:grid-cols-3 md:gap-6">
        <div className="md:col-span-1">
          <div className="px-4 sm:px-0">
            <h3 className="text-lg font-medium leading-6 text-gray-900">Sign Up</h3>
            <p className="mt-1 text-sm text-gray-600"> Please type name which you want to use. </p>
          </div>
        </div>
        <div className="mt-5 md:col-span-2 md:mt-0">
          <form action="#" method="POST">
            <div className="overflow-hidden shadow sm:rounded-md">
              <div className="bg-white px-4 py-5 sm:p-6">
                <div className="grid grid-cols-6 gap-6">
                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="user_name" className="block text-sm font-medium text-gray-700">User Name</label>
                    <input type="text" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm tascholaForm" value={userName} onChange={(event) => setUserName(event?.target.value)} />
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="password1" className="block text-sm font-medium text-gray-700">Password</label>
                    <input type="password" name="password1" id="password1" autoComplete="password" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm tascholaForm" value={password1} onChange={(event) => setPassword1(event?.target.value)} />
                  </div>

                  <div className="col-span-6 sm:col-span-4">
                    <label htmlFor="password2" className="block text-sm font-medium text-gray-700">Password (for confirmed)</label>
                    <input type="password" name="password2" id="password2" autoComplete="password" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm tascholaForm" value={password2} onChange={(event) => setPassword2(event?.target.value)} />
                  </div>

                  <div className="col-span-6 sm:col-span-3">
                    <label htmlFor="student_type" className="block text-sm font-medium text-gray-700">Student Type</label>
                    <select id="student_type" name="student_type" autoComplete="degree" className="mt-1 block w-full rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm tascholaForm">
                      <option>Bachelor</option>
                      <option>Master</option>
                      <option>Doctor</option>
                    </select>
                  </div>
                </div>
              </div>
              <div className="bg-gray-50 px-4 py-3 text-right sm:px-6">
                <button type="submit" className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2" onClick={handleSubmit}>Sign Up</button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}

export default UserRegisterForm
