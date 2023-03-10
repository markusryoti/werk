import { useRouter } from 'next/router'
import React, { useEffect } from 'react'
import { useAuth } from '../../context/AuthUserContext'

const Profile = () => {
    const router = useRouter()
    const { authUser, logout, loading } = useAuth()

    useEffect(() => {
        if (!authUser && !loading) {
            router.push("/login")
        }
    }, [authUser, router, loading])

    const doLogout = async () => {
        await logout()
        router.push("/login")
    }

    return (
        <div className="flex justify-center container mx-auto pt-8 p-1">
            <button onClick={doLogout} className="btn btn-warning w-full md:w-1/2">
                Logout
            </button>
        </div>
    )
}

export default Profile
