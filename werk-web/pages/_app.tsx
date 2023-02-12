import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Nav from '../components/nav'
import Footer from '../components/footer'
import { AuthContextProvider } from '../context/AuthUserContext'


export default function App({ Component, pageProps }: AppProps) {
    return (
        <AuthContextProvider>
            <Nav />
            <Component {...pageProps} />
            <Footer />
        </AuthContextProvider >
    )
}
