import LoginForm from '@/components/auth/LoginForm'
import styles from '@/styles/Home.module.scss'

const Login = () => {
  return (
    <>
      <main className={styles.main}>
        <LoginForm />
      </main>
    </>
  )
}

export default Login
