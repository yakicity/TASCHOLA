import { User } from '@/interfaces/user'
import styles from '@/styles/Home.module.scss'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import Link from 'next/link'
import { useEffect, useState } from 'react'

const UserInfo = () => {
  useEffect(() => {
    try {
      axios.get(`${url}/v1/user/2`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        }
      }) // TODO: get user id from session
        .then((res: AxiosResponse<User>) => {
          const { data, status } = res
          switch (status) {
            case 200:
              setUser(data)
              break
            case 404:
              alert('User not found' + res.statusText)
              break
            default:
              alert('Something went wrong' + res.statusText)
          }
        })
    } catch (error) {
      alert(console.error())
    }
  }, [])

  const [user, setUser] = useState<User>()

  return (
    <>
      <main className={styles.main}>
        <h1>User Info</h1>
        {user && (
          <div>
            <div>
              <span>Id: </span>
              <span>{user.id}</span>
            </div>
            <div>
              <span>Username: </span>
              <span>{user.name}</span>
            </div>
            <div>
              <Link href={`/user/edit/`} className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">
                Edit User
              </Link>
            </div>
          </div>
        )}
      </main>
    </>
  )
}

export default UserInfo
