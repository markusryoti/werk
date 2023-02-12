import { createContext, useContext, useEffect, useState } from 'react'
import {
    onAuthStateChanged,
    createUserWithEmailAndPassword,
    signInWithEmailAndPassword,
    signOut,
    UserCredential,
} from 'firebase/auth'
import { auth } from '../config/firebase';

export interface AuthContext {
    authUser: AuthUser | null;
    login: (email: string, password: string) => Promise<void>
    signup: (email: string, password: string) => Promise<UserCredential>
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
                console.log('auth state changed, has user')

                setAuthUser({
                    uid: user.uid,
                    email: user.email,
                    displayName: user.displayName
                })

                console.log('starting a session')
                const token = await user.getIdToken()
                await doSessionLogin(token)
            } else {
                console.log('resetting auth user')
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

    const signup = (email: string, password: string): Promise<UserCredential> => {
        setLoading(true)
        return createUserWithEmailAndPassword(auth, email, password)
    }

    const login = async (email: string, password: string) => {
        setLoading(true)

        await signInWithEmailAndPassword(auth, email, password)
    }

    const logout = async () => {
        await signOut(auth)
        await doSessionLogout()
    }

    const getToken = async (): Promise<string | undefined> => {
        return auth.currentUser?.getIdToken(true)
    }

    return (
        <authContext.Provider value={{ authUser, login, signup, logout, getToken, loading }}>
            {loading ? null : children}
        </authContext.Provider>
    )
}
