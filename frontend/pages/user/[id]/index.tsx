import styles from '@/styles/Home.module.scss'
import { useRouter } from 'next/router'

const UserInfo = () => {
  const router = useRouter()
  const { id } = router.query

  return (
    <>
      <main className={styles.main}>
        <h1>User Info</h1>
        <p>id: {id}</p>
      </main>
    </>
  )
}

export default UserInfo
