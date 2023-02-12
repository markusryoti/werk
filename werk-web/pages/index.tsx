import Head from 'next/head'
import { useRouter } from 'next/router'
import { useAuth } from '../context/AuthUserContext'
import styles from '../styles/Home.module.css'

export default function Home() {
    const router = useRouter()
    const { authUser, loading } = useAuth()

    if (authUser && !loading) {
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
