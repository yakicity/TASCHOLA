import styles from '@/styles/Home.module.scss'
import { useRouter } from 'next/router'

const TaskPage = () => {
  const router = useRouter()
  const { id } = router.query

  return (
    <>
      <main className={styles.main}>
        <h1>Task {id}</h1>
      </main>
    </>
  )
}

export default TaskPage
