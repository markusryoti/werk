import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import { useClientRequest } from "../../../hooks/use-request"
import { Movement } from "../../types"

import { CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts";

interface MovementStats {
    movement: Movement;
    recentDevelopment: RecentMovementDevelopment;
    estimatedMaxes: EstimatedMax[]
}

interface RecentMovementDevelopment {
    change: number;
    numWorkouts: number;
}

interface EstimatedMax {
    x: Date;
    y: number;
}

export default function MovementAnalytics() {
    const [stats, setStats] = useState({} as MovementStats)

    const router = useRouter()
    const { id } = router.query

    const { doRequest } = useClientRequest()

    useEffect(() => {
        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/movements/${id}`
        doRequest(url, 'GET')
            .then(res => res.json())
            .then(s => setStats(s))
            .catch(err => console.error(err))
    }, [])

    useEffect(() => {
        if (stats.estimatedMaxes && stats.estimatedMaxes.length > 0) {
            const labels = [];
            const values = [];

            for (let point of stats.estimatedMaxes) {
                const { x, y } = point;
                labels.push(x);
                values.push(y);
            }
        }
    }, [stats])

    return (
        <div className="container mx-auto pt-8 p-1">
            <ResponsiveContainer width="90%" aspect={3}>
                <LineChart data={stats.estimatedMaxes} margin={{ right: 20, left: -20 }}>
                    <CartesianGrid stroke="#243240" />
                    <XAxis dataKey="x" />
                    <YAxis />
                    <Legend verticalAlign="top" height={36} />
                    <Line type="monotone" dataKey="y" stroke="#82ca9d" strokeWidth={3} />
                </LineChart>
            </ResponsiveContainer>
        </div >
    )
}
