import { useRouter } from 'next/router'
import React from 'react'
import { useAuth } from '../../context/AuthUserContext'

const Profile = () => {
    const router = useRouter()
    const { authUser, logout } = useAuth()

    if (!authUser) {
        router.push("/login")
    }

    const doLogout = async () => {
        await logout()
        router.push("/login")
    }

    return (
        <div className="container mx-auto pt-8 p-1">
            <button onClick={doLogout} className="btn btn-warning w-full md:w-1/2">
                Logout
            </button>
        </div>
    )
}

export default Profile
