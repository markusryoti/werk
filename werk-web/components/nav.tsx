import Link from "next/link"
import { useAuth } from "../pages/context/AuthUserContext"

export default function Nav() {
    const { authUser } = useAuth()
    return (
        <div className="navbar bg-primary text-primary-content">
            <div className="flex-1">
                <Link href="/workouts" className="btn btn-ghost normal-case text-xl">werk</Link>
            </div>
            {authUser && (
                <div className="flex-1 justify-end">
                    <Link href="/profile" className="btn btn-ghost">
                        {authUser.email}
                    </Link>
                </div>
            )}
        </div>
    )
}
