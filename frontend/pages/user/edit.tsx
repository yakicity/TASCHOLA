import { UpdateUserForm, User, UserForm } from '@/interfaces/user'
import styles from '@/styles/Home.module.scss'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import Router from 'next/router'
import { useEffect, useState } from 'react'
import Cookies from 'universal-cookie'

const UserEditPage = () => {
  useEffect(() => {
    try {
      // get user_id from cookie
      const cookies = new Cookies()
      const userID = cookies.get('user_id')
      setRequestUserID(userID)
      if (!userID) {
        Router.push('/login')
        return
      }

      axios.get(`${url}/v1/user/${userID}`, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then((res: AxiosResponse<User>) => {
          const { data, status } = res
          switch (status) {
            case 200:
              setUser(data)
              setLoaded(true)
              break
            case 404:
              alert('User not found')
              break
            default:
              alert('Something went wrong')
          }
        })
    } catch (error) {
      alert(console.error())
    }
  }, [])

  const [user, setUser] = useState<User>()
  const [loaded, setLoaded] = useState<boolean>(false)

  const [name, setName] = useState<string>('')
  const [oldPassword, setOldPassword] = useState<string>('')
  const [newPassword, setNewPassword] = useState<string>('')
  const [requestUserID, setRequestUserID] = useState<string>('')

  const handleSubmit = () => {
    const updateUserForm: UpdateUserForm = {
      old_password: oldPassword,
      user: {
        name: name,
        password: newPassword
      } as UserForm
    }

    axios.put(`${url}/v1/user/${requestUserID}`, updateUserForm, {
      headers: {
        'Content-Type': 'application/json'
      }
    }).then((res => {
      const { status } = res
      switch (status) {
        case 200:
          alert('User updated successfully')
          break
        default:
          alert('Something went wrong' + res.statusText)
      }
    }))
  }

  return (
    <>
      <main className={styles.main}>
        {!loaded && <div>Loading...</div>}
        {loaded && !user && <div>No user found</div>}
        {loaded && user && (
          <div className="mt-10 sm:mt-0">
            <div className="md:grid md:grid-cols-3 md:gap-6">
              <div className="md:col-span-1">
                <div className="px-4 sm:px-0">
                  <h3 className="text-lg font-medium leading-6 text-gray-900">Change Your User Info </h3>
                </div>
              </div>
              <div className="mt-5 md:col-span-2 md:mt-0">
                <form action="#" method="POST">
                  <div className="overflow-hidden shadow sm:rounded-md">
                    <div className="bg-white px-4 py-5 sm:p-6">
                      <div className="grid grid-cols-6 gap-6">
                        <div className="col-span-6 sm:col-span-4">
                          <label htmlFor="user_name" className="block text-sm font-medium text-gray-700">New User Name(Don't duplicate)</label>
                          <input type="text" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm tascholaForm" value={name} onChange={(event) => setName(event?.target.value)} />
                        </div>
                        <div className="col-span-6 sm:col-span-4">
                          <label className="block text-sm font-medium text-gray-700">Old Password</label>
                          <input type="password" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm tascholaForm" value={oldPassword} onChange={(event) => setOldPassword(event?.target.value)} />
                        </div>
                        <div className="col-span-6 sm:col-span-4">
                          <label className="block text-sm font-medium text-gray-700">New Password</label>
                          <input type="password" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm tascholaForm" value={newPassword} onChange={(event) => setNewPassword(event?.target.value)} />
                        </div>
                      </div>
                    </div>
                    <div className="bg-gray-50 px-4 py-3 text-right sm:px-6">
                      <button type="submit" className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2" onClick={handleSubmit}>Change</button>
                    </div>
                  </div>
                </form>
              </div>
            </div>
          </div>
        )}
      </main>
    </>
  )
}

export default UserEditPage
