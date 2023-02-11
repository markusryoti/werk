import Head from 'next/head'
import { useRouter } from 'next/router'
import styles from '../styles/Home.module.css'
import { useAuth } from './context/AuthUserContext'

export default function Home() {
    const router = useRouter()
    const { authUser } = useAuth()

    if (authUser) {
        router.push("/workouts")
    }

    return (
        <>
            <Head>
                <title>Werk</title>
            </Head>
            <main className={styles.main}>
                Empty page
            </main>
        </>
    )
}
