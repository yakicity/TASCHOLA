import { useRouter } from 'next/router'

const TaskPage = () => {
  const router = useRouter()
  const { id } = router.query

  return (
    <div>
      <h1>Task {id}</h1>
    </div>
  )
}

export default TaskPage
