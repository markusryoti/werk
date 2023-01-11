import React, { FormEvent, useState } from 'react'
import { useAuth } from '../context/AuthUserContext'

export default function Login() {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const { login } = useAuth()

    const onSubmit = (e: FormEvent) => {
        e.preventDefault()

        if (!email || !password) {
            return
        }

        login(email, password)
            .then(res => console.log(res))
            .catch(err => console.log(err))
    }

    return (
        <div className="flex justify-center">
            <div className="card w-1/2 bg-base-300 shadow-xl mt-16">
                <form onSubmit={onSubmit} className="p-8">
                    <div className="flex m-2">
                        <label htmlFor="email" className="flex-1">Email</label>
                        <input id="email" type="text" onChange={e => setEmail(e.target.value.trim())} className="input w-full" />
                    </div>
                    <div className="flex m-2">
                        <label htmlFor="password" className="flex-1">Password</label>
                        <input id="password" type="password" onChange={e => setPassword(e.target.value.trim())} className="input w-full" />
                    </div>
                    <div className="flex justify-center">
                        <div className="mt-8">
                            <button type="submit" className="btn btn-primary">Login</button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    )
}
