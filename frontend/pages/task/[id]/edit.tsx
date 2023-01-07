import { useRouter } from 'next/router'

const TaskEditPage = () => {
  const router = useRouter()
  const { id } = router.query

  return (
    <div>
      <h1>Edit Task {id}</h1>
    </div>
  )
}

export default TaskEditPage
