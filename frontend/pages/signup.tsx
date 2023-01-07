import UserRegisterForm from "@/components/UserRegisterForm"
import styles from '@/styles/Home.module.scss'

const SignUp = () => {
  return (
    <>
      <main className={styles.main}>
        <UserRegisterForm />
      </main>
    </>
  )
}

export default SignUp
