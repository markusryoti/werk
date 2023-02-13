import Head from 'next/head'
import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAuth } from '../context/AuthUserContext'
import styles from '../styles/Home.module.css'

export default function Home() {
    const router = useRouter()
    const { authUser, loading } = useAuth()

    useEffect(() => {
        if (authUser && !loading) {
            router.push("/workouts")
        }
    }, [authUser, loading, router])

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
