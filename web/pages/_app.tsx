import { ChakraProvider, extendTheme } from '@chakra-ui/react'
import type { AppProps } from 'next/app'
import '../styles/fonts.css'
import '../styles/youtube.css'
import Head from 'next/head'
import { Provider } from 'react-redux'
import { store } from '@/redux/store'
import NextNProgress from 'nextjs-progressbar';
import '@/styles/blog.css'
import Script from 'next/script'

const theme = extendTheme({
  fonts: {
    heading: `'Rubik', sans-serif`,
    body: `'Rubik', sans-serif`,
  },
})

export default function App({ Component, pageProps }: AppProps) {
  return (
    <Provider store={store}>
      <NextNProgress color="#38a169" />
      <ChakraProvider theme={theme}>
        <Head>
          <meta name="application-name" content="Manejemen Produk" />
          <meta name="apple-mobile-web-app-capable" content="yes" />
          <meta name="apple-mobile-web-app-status-bar-style" content="default" />
          <meta name="apple-mobile-web-app-title" content="Manajemen Produk" />
          <meta name="description" content="Best PWA App in the world" />
          <meta name="format-detection" content="telephone=no" />
          <meta name="mobile-web-app-capable" content="yes" />
          <meta name="msapplication-config" content="/generated_icon/browserconfig.xml" />
          <meta name="msapplication-tap-highlight" content="no" />

          <link rel="icon" href="/favicon.ico" />
          <link rel="apple-touch-icon" sizes="57x57" href="/generated_icon/apple-icon-57x57.png" />
          <link rel="apple-touch-icon" sizes="60x60" href="/generated_icon/apple-icon-60x60.png" />
          <link rel="apple-touch-icon" sizes="72x72" href="/generated_icon/apple-icon-72x72.png" />
          <link rel="apple-touch-icon" sizes="76x76" href="/generated_icon/apple-icon-76x76.png" />
          <link rel="apple-touch-icon" sizes="114x114" href="/generated_icon/apple-icon-114x114.png" />
          <link rel="apple-touch-icon" sizes="120x120" href="/generated_icon/apple-icon-120x120.png" />
          <link rel="apple-touch-icon" sizes="144x144" href="/generated_icon/apple-icon-144x144.png" />
          <link rel="apple-touch-icon" sizes="152x152" href="/generated_icon/apple-icon-152x152.png" />
          <link rel="apple-touch-icon" sizes="180x180" href="/generated_icon/apple-icon-180x180.png" />
          <link rel="icon" type="image/png" sizes="192x192" href="/generated_icon/android-icon-192x192.png" />
          <link rel="icon" type="image/png" sizes="144x144" href="/generated_icon/android-icon-144x144.png" />
          <link rel="icon" type="image/png" sizes="32x32" href="/generated_icon/favicon-32x32.png" />
          <link rel="icon" type="image/png" sizes="96x96" href="/generated_icon/favicon-96x96.png" />
          <link rel="icon" type="image/png" sizes="16x16" href="/generated_icon/favicon-16x16.png" />
          <link rel="manifest" href="/generated_icon/manifest.json" />
          <meta name="msapplication-TileColor" content="#38a169" />
          <meta name="msapplication-TileImage" content="/generated_icon/ms-icon-144x144.png" />
          <meta name="theme-color" content="#38a169" />

          <meta
            name='viewport'
            content='minimum-scale=1, initial-scale=1, width=device-width, shrink-to-fit=no, viewport-fit=cover'
          />
        </Head>
        <Script async src="https://www.googletagmanager.com/gtag/js?id=G-S5QSJVVSZQ" strategy='beforeInteractive' />
        <Script src={"/google_analytic.js"} strategy='beforeInteractive' />
        <Component {...pageProps} />
      </ChakraProvider>
    </Provider>
  )
}
