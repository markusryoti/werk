import Link from "next/link"
import { useEffect, useState } from "react"
import { useClientRequest } from "../../hooks/use-request"
import { Movement } from "../types"

export default function Analytics() {
    const [movements, setMovements] = useState<Movement[]>([])

    const { doRequest } = useClientRequest()

    useEffect(() => {
        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/movements`

        doRequest(url, 'GET')
            .then(res => res.json())
            .then(m => setMovements(m))
            .catch(err => console.error(err))
    }, [])

    return (
        <div className="container mx-auto pt-8 p-1">
            {movements.length > 0 && movements.map((movement: Movement) => {
                return <div key={movement.id}>
                    <Link href={`/analytics/movements/${movement.id}`}>
                        {movement.name}
                    </Link>
                </div>
            })}
        </div >
    )
}
