import styles from '@/styles/Home.module.scss'
import NewTaskForm from '@/components/task/NewTaskForm'

const TaskNewPage = () => {
  return (
    <>
      <main className={styles.main}>
        <NewTaskForm />
      </main>
    </>
  )
}

export default TaskNewPage
