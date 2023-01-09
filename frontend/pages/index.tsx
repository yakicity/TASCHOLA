import { Inter } from '@next/font/google'
import Head from 'next/head'
import Image from 'next/image'
import styles from 'styles/Home.module.scss'

const inter = Inter({ subsets: ['latin'] })

const Home = () => {
  return (
    <>
      <Head>
        <title>TASCHOLA</title>
        <meta name="description" content="TASCHOLA is task management tool specific to Tokyo Institute of Technology students." />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={styles.main}>
        <div className={styles.description}>
          <p>
            TASCOLA make it easy to manage your tasks.
          </p>
          <div>
            <a
              href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
              target="_blank"
              rel="noopener noreferrer"
            >
              By{' '}
              <Image
                src="/vercel.svg"
                alt="Vercel Logo"
                className={styles.vercelLogo}
                width={100}
                height={24}
                priority
              />
            </a>
          </div>
        </div>

        <div className={styles.center}>
          <Image
            className={styles.logo}
            src="/taschola.svg"
            alt="TASCHOLA Logo"
            width={280}
            height={37}
            priority
          />
        </div>

        <div className={styles.grid}>
          <a
            href="/signup"
            className={styles.card}
            rel="noopener noreferrer"
          >
            <h2 className={inter.className}>
              Sign Up <span>-&gt;</span>
            </h2>
            <p className={inter.className}>
              Sign up for a new account.
            </p>
          </a>

          <a
            href="/login"
            className={styles.card}
            rel="noopener noreferrer"
          >
            <h2 className={inter.className}>
              Login <span>-&gt;</span>
            </h2>
            <p className={inter.className}>
              Login to your account.
            </p>
          </a>

          <a href="/tasks" className={styles.card} rel="noopener noreferrer">
            <h2 className={inter.className}>
              Your Tasks <span>-&gt;</span>
            </h2>
            <p className={inter.className}>
              View your tasks.
            </p>
          </a>

          <a
            href="/user"
            className={styles.card}
            rel="noopener noreferrer"
          >
            <h2 className={inter.className}>
              User Info <span>-&gt;</span>
            </h2>
            <p className={inter.className}>
              View your user info.
              Only for logged in users
            </p>
          </a>
        </div>
      </main>
    </>
  )
}

export default Home
