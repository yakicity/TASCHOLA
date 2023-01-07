import UserRegisterForm from "@/components/UserRegisterForm"
import styles from '@/styles/Home.module.scss'

const SignUp = () => {
  return (
    <>
      <main className={styles.main}>
        <h1>SignUp</h1>
        <UserRegisterForm />
      </main>
    </>
  )
}

export default SignUp
