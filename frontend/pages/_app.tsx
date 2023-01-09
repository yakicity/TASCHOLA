import Layout from '@/components/layout/layout'
import type { AppProps } from 'next/app'
import 'styles/globals.scss'

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <Layout>
      <Component {...pageProps} />
    </Layout>
  )
}

export default App
