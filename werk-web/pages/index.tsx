import Head from 'next/head'
import styles from '../styles/Home.module.css'
import { useAuth } from './context/AuthUserContext'

export default function Home() {
    const { authUser } = useAuth()
    return (
        <>
            <Head>
                <title>Werk</title>
            </Head>
            <main className={styles.main}>
                <pre>{JSON.stringify(authUser)}</pre>
            </main>
        </>
    )
}
