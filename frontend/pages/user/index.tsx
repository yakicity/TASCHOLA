import { User } from '@/interfaces/user'
import styles from '@/styles/Home.module.scss'
import { url } from '@/utils/constants'
import axios, { AxiosResponse } from 'axios'
import Link from 'next/link'
import Router from 'next/router'
import { useEffect, useState } from 'react'
import { AiOutlineUser } from 'react-icons/ai'
import Cookies from 'universal-cookie'

const UserInfo = () => {
  useEffect(() => {
    // get user_id from cookie
    const cookies = new Cookies()
    const userID = cookies.get('user_id')
    if (!userID) {
      Router.push('/login')
      return
    }

    // API Request
    try {
      axios.get(`${url}/v1/user/${userID}`, {})
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
        {user && (
          <div>
            <section className="flex font-medium items-center justify-center h-screen px-6 rounded-xl">

              <section className="w-64 mx-auto bg-[#6f8ba9] rounded-2xl px-8 py-6 shadow-lg">
                <div className="flex items-center justify-between">
                  <span className="text-emerald-400">
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z" />
                    </svg>
                  </span>
                </div>
                <div className="mt-6 w-fit mx-auto">
                  <AiOutlineUser className="" />
                </div>
                <div className="mt-8 ">
                  <h2 className="text-white font-bold text-2xl tracking-wide">{user.name}</h2>
                </div>
                <p className="text-emerald-400 font-semibold mt-2.5" >
                  id: {user.id}
                </p>

                <div className="mt-10">
                  <Link href={`/user/edit/`} className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">
                    Edit User
                  </Link>
                </div>
              </section>
            </section>
          </div>
        )}
      </main>
    </>
  )
}

export default UserInfo
