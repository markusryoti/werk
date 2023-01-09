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
    user: AuthUser | null;
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
    const [user, setUser] = useState<AuthUser | null>(null)
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        const unsubscribe = onAuthStateChanged(auth, (user) => {
            if (user) {
                setUser({
                    uid: user.uid,
                    email: user.email,
                    displayName: user.displayName
                })
            } else {
                setUser(null)
            }
            setLoading(false)
        })

        return () => unsubscribe()
    }, [])

    const signup = (email: string, password: string): Promise<UserCredential> => {
        return createUserWithEmailAndPassword(auth, email, password)
    }

    const login = (email: string, password: string): Promise<UserCredential> => {
        return signInWithEmailAndPassword(auth, email, password)
    }

    const logout = async () => {
        setUser(null)
        await signOut(auth)
    }

    const getToken = async (): Promise<string | undefined> => {
        return auth.currentUser?.getIdToken(true)
    }

    return (
        <authContext.Provider value={{ user, login, signup, logout, getToken }}>
            {loading ? null : children}
        </authContext.Provider>
    )
}
