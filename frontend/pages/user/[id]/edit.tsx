import { useRouter } from 'next/router'

const UserEditPage = () => {
  const router = useRouter()
  const { id } = router.query

  return (
    <div>
      <h1>Edit User {id}</h1>
    </div>
  )
}

export default UserEditPage
