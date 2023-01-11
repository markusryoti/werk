import { createContext, useContext, useEffect, useState } from 'react'
import { auth } from '../../config/firebase'
import {
    onAuthStateChanged,
    createUserWithEmailAndPassword,
    signInWithEmailAndPassword,
    signOut,
    UserCredential,
} from 'firebase/auth'

export interface AuthContext {
    authUser: AuthUser | null;
    login: (email: string, password: string) => Promise<UserCredential>
    signup: (email: string, password: string) => Promise<UserCredential>
    logout: () => Promise<void>
    getToken: () => Promise<string | undefined>
}

export type AuthUser = {
    uid: string;
    email: string | null
    displayName: string | null
}

const authContext = createContext<AuthContext>({} as AuthContext)

export const useAuth = () => useContext(authContext)

export const AuthContextProvider = ({
    children,
}: {
    children: React.ReactNode
}) => {
    const [authUser, setAuthUser] = useState<AuthUser | null>({} as AuthUser)
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        const unsubscribe = onAuthStateChanged(auth, async (user) => {
            if (user) {
                setAuthUser({
                    uid: user.uid,
                    email: user.email,
                    displayName: user.displayName
                })

                const token = await user.getIdToken()
                await doSessionLogin(token)
            } else {
                setAuthUser(null)
                await doSessionLogout()
            }
            setLoading(false)
        })

        return () => {
            unsubscribe()
        }
    }, [])

    const doSessionLogin = async (token: string) => {
        const res = await fetch('/api/session', {
            method: 'post',
            body: JSON.stringify({ token }),
        })

        const body = await res.json()
        console.log(body)
    }

    const doSessionLogout = async () => {
        await fetch('/api/logout')
    }

    const signup = (email: string, password: string): Promise<UserCredential> => {
        return createUserWithEmailAndPassword(auth, email, password)
    }

    const login = (email: string, password: string): Promise<UserCredential> => {
        return signInWithEmailAndPassword(auth, email, password)
    }

    const logout = async () => {
        setAuthUser(null)
        await signOut(auth)
    }

    const getToken = async (): Promise<string | undefined> => {
        return auth.currentUser?.getIdToken(true)
    }

    return (
        <authContext.Provider value={{ authUser, login, signup, logout, getToken }}>
            {loading ? null : children}
        </authContext.Provider>
    )
}
