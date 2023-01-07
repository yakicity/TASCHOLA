import styles from '@/styles/Home.module.scss'
import { useRouter } from 'next/router'

const TaskEditPage = () => {
  const router = useRouter()
  const { id } = router.query

  return (
    <>
      <main className={styles.main}>
        <h1>Edit Task {id}</h1>
      </main>
    </>
  )
}

export default TaskEditPage
