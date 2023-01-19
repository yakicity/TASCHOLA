import { User, UserForm } from '@/interfaces/user'
import styles from '@/styles/Home.module.scss'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import { useEffect, useState } from 'react'

const UserEditPage = () => {
  useEffect(() => {
    try {
      axios.get(`${url}/v1/users/1`) // TODO: get user id from session
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
  const [password, setPassword] = useState<string>('')

  const handleSubmit = () => {
    const userForm: UserForm = {
      name: name,
      password: password
    }

    axios.put(`${url}/v1/users/1`, userForm)
      .then((res => {
        const { status } = res
        switch (status) {
          case 200:
            alert('User updated successfully')
            break
          default:
            alert('Something went wrong')
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
                  <h3 className="text-lg font-medium leading-6 text-gray-900">Edit Your task Task</h3>
                </div>
              </div>
              <div className="mt-5 md:col-span-2 md:mt-0">
                <form action="#" method="POST">
                  <div className="overflow-hidden shadow sm:rounded-md">
                    <div className="bg-white px-4 py-5 sm:p-6">
                      <div className="grid grid-cols-6 gap-6">
                        <div className="col-span-6 sm:col-span-4">
                          <label htmlFor="user_name" className="block text-sm font-medium text-gray-700">User Name</label>
                          <input type="text" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" value={name} onChange={(event) => setName(event?.target.value)} />
                        </div>
                        <div className="col-span-6 sm:col-span-4">
                          <label htmlFor="user_password" className="block text-sm font-medium text-gray-700">User Password</label>
                          <input type="text" className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm" value={""} onChange={(event) => setPassword(event?.target.value)} />
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
