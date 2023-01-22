import Layout from '@/components/layout/Layout'
import '@/styles/globals.scss'
import { url } from '@/utils/constants'
import axios from 'axios'
import type { AppProps } from 'next/app'

const App = ({ Component, pageProps }: AppProps) => {
  // axios setting
  axios.defaults.withCredentials = true
  axios.defaults.baseURL = url

  return (
    <Layout>
      <Component {...pageProps} />
    </Layout>
  )
}

export default App
