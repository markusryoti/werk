import Link from "next/link"
import { useAuth } from "../pages/context/AuthUserContext"

export default function Nav() {
    const { authUser } = useAuth()
    return (
        <div className="navbar bg-base-300">
            <div className="flex-none">
                <button className="btn btn-square btn-ghost">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" className="inline-block w-5 h-5 stroke-current"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h16M4 18h16"></path></svg>
                </button>
            </div>
            <div className="flex-1">
                <Link href="/workouts" className="btn btn-ghost normal-case text-xl">Werk</Link>
            </div>

            {authUser && (
                <div className="flex-1 justify-end">
                    <button className="btn btn-ghost">
                        {authUser.email}
                    </button>
                </div>
            )}
        </div>
    )
}
