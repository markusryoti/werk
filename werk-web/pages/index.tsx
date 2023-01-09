import Head from 'next/head'
import styles from '../styles/Home.module.css'
import { useAuth } from './context/AuthUserContext'

export default function Home() {
    const { user } = useAuth()
    return (
        <>
            <Head>
                <title>Werk</title>
            </Head>
            <main className={styles.main}>
                <h1 className="text-3xl font-bold underline">Werk</h1>
                <pre>{JSON.stringify(user)}</pre>
            </main>
        </>
    )
}
