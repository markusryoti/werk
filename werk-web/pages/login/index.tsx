import React, { FormEvent, useState } from 'react'
import { useRouter } from 'next/router'
import { useAuth } from '../../context/AuthUserContext'

export default function Login() {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const router = useRouter()
    const { authUser, login, loading } = useAuth()

    const onSubmit = async (e: FormEvent) => {
        e.preventDefault()

        if (!email || !password) {
            return
        }

        await login(email, password)
        router.push("/workouts")
    }

    if (authUser && !loading) {
        router.push("/workouts")
    }

    return (
        <div className="flex justify-center">
            <div className="card w-full md:w-1/2 bg-base-300 shadow-xl mt-16 m-2">
                <h1 className='text-xl p-2'>Let&apos;s get started</h1>
                <form onSubmit={onSubmit} className="p-4">
                    <div className="flex flex-col p-1">
                        <label htmlFor="email" className="">Email</label>
                        <input id="email" type="text" onChange={e => setEmail(e.target.value.trim())} className="input" />
                    </div>
                    <div className="flex flex-col p-1">
                        <label htmlFor="password" className="">Password</label>
                        <input id="password" type="password" onChange={e => setPassword(e.target.value.trim())} className="input" />
                    </div>
                    <div className="flex justify-center">
                        <div className="mt-6">
                            <button type="submit" className="btn btn-primary">Login</button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    )
}
