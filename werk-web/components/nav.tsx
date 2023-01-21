import Link from "next/link"
import { useAuth } from "../pages/context/AuthUserContext"

export default function Nav() {
    const { authUser } = useAuth()
    return (
        <div className="navbar bg-base-300">
            <div className="flex-1">
                <Link href="/workouts" className="btn btn-ghost normal-case text-xl">werk</Link>
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
