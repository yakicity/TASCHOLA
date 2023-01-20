import Layout from '@/components/layout/Layout'
import type { AppProps } from 'next/app'
import { useRouter } from 'next/router'
import { url } from '@/utils/constants'
import 'styles/globals.scss'
import axios from 'axios'
import { useEffect, useState } from 'react'

const App = ({ Component, pageProps }: AppProps) => {
  // axios setting
  axios.defaults.withCredentials = true
  axios.defaults.baseURL = url

  const router = useRouter()
  // tasks/*, user/*, is required to login
  // jwt token is required to access to these pages
  if (router.pathname.startsWith('/tasks') || router.pathname.startsWith('/user')) {

  }
  return (
    <Layout>
      <Component {...pageProps} />
    </Layout>
  )
}

export default App
