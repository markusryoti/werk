import { createContext, useContext, useEffect, useState } from 'react'
import {
    onAuthStateChanged,
    createUserWithEmailAndPassword,
    signInWithEmailAndPassword,
    signOut,
} from 'firebase/auth'
import { auth } from '../config/firebase';

export interface AuthContext {
    authUser: AuthUser | null;
    login: (email: string, password: string) => Promise<void>
    signup: (email: string, password: string) => Promise<void>
    logout: () => Promise<void>
    getToken: () => Promise<string | undefined>
    loading: boolean
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
            } else {
                setAuthUser(null)
                await doSessionLogout()
            }
        })

        return () => {
            unsubscribe()
        }
    }, [])

    const doSessionLogin = async (token: string) => {
        setLoading(true)

        await fetch('/api/session', {
            method: 'post',
            body: JSON.stringify({ token }),
        })
    }

    const doSessionLogout = async () => {
        setLoading(true)
        await fetch('/api/logout')
    }

    const signup = async (email: string, password: string): Promise<void> => {
        setLoading(true)

        const userCreds = await createUserWithEmailAndPassword(auth, email, password)
        const token = await userCreds.user.getIdToken(true)
        await doSessionLogin(token)

        setLoading(false)
    }

    const login = async (email: string, password: string) => {
        setLoading(true)

        const userCreds = await signInWithEmailAndPassword(auth, email, password)
        const token = await userCreds.user.getIdToken(true)
        await doSessionLogin(token)

        setLoading(false)
    }

    const logout = async () => {
        await signOut(auth)
        await doSessionLogout()
    }

    const getToken = async (): Promise<string | undefined> => {
        const token = await auth.currentUser?.getIdToken(true)
        await doSessionLogin(token!)
        return token
    }

    return (
        <authContext.Provider value={{ authUser, login, signup, logout, getToken, loading }}>
            {children}
        </authContext.Provider>
    )
}
