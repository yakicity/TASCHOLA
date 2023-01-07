import { useRouter } from 'next/router'

const UserInfo = () => {
  const router = useRouter()
  const { id } = router.query

  return (
    <div>
      <h1>User Info</h1>
      <p>id: {id}</p>
    </div>
  )
}

export default UserInfo
