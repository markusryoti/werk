import Link from "next/link"
import { useEffect, useState } from "react"
import Spinner from "../../components/spinner"
import { useClientRequest } from "../../hooks/use-request"
import { Movement } from "../../lib/types"

export default function Analytics() {
    const [movements, setMovements] = useState<Movement[]>([])

    const { doRequest, waiting } = useClientRequest()

    useEffect(() => {
        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/movements`

        doRequest(url, 'GET')
            .then(res => res.json())
            .then(m => setMovements(m))
            .catch(err => console.error(err))
    }, [doRequest])

    if (waiting) {
        return <Spinner />
    }

    return (
        <div className="flex justify-center container mx-auto pt-8 p-1">
            <div className="w-full md:w-1/2">
                {movements.length > 0 && movements.map((movement: Movement) => {
                    return <div key={movement.id} className="card border border-base-300 bg-base-200 mt-2 p-3">
                        <Link href={`/analytics/movements/${movement.id}`} className="card-title">
                            {movement.name}
                        </Link>
                    </div>
                })}
            </div>
        </div >
    )
}
