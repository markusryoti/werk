import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import { useClientRequest } from "../../../hooks/use-request"

import { CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts";
import Spinner from "../../../components/spinner";
import { Movement } from "../../../lib/types";

type MovementStats = {
    movement: Movement;
    recentDevelopment: RecentMovementDevelopment;
    estimatedMaxes: EstimatedMax[]
    currentMax: EstimatedMax
    allTimeMax: EstimatedMax
}

type RecentMovementDevelopment = {
    change: number;
    numWorkouts: number;
}

type EstimatedMax = {
    date: string | Date;
    max: number;
}

export default function MovementAnalytics() {
    const [stats, setStats] = useState<MovementStats | undefined>(undefined)

    const router = useRouter()
    const { id } = router.query

    const { doRequest, waiting } = useClientRequest()

    useEffect(() => {
        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/movements/${id}`
        doRequest(url, 'GET')
            .then(res => res.json())
            .then(s => {
                const chartData = s.estimatedMaxes.map((point: { date: string, max: number }) => {
                    const { date, max } = point
                    return {
                        date: new Date(date).toLocaleDateString(),
                        max: max.toFixed(2)
                    }
                })

                setStats({ ...s, estimatedMaxes: chartData })
            })
            .catch(err => console.error(err))
        // eslint-disable-next-line
    }, [])

    if (waiting) {
        return <Spinner />
    }

    if (!stats || stats.estimatedMaxes?.length === 0) {
        return (
            <div className="container mx-auto pt-8 p-1">
                <h1 className="text-2xl mb-4">No maxes to report</h1>
                <p>Add movement sets and come back</p>
            </div >
        )
    }

    const getChartHeight = () => {
        return window.innerHeight * 0.5
    }

    return (
        <div className="container mx-auto pt-8 p-2">
            <h1 className="text-2xl ml-4">{stats.movement.name}</h1>
            <ResponsiveContainer width="100%" height={getChartHeight()}>
                <LineChart data={stats?.estimatedMaxes} margin={{ top: 10, right: 60, left: 0, bottom: 0 }}>
                    <CartesianGrid stroke="#243240" />
                    <XAxis dataKey="date" />
                    <YAxis type="number" domain={['dataMin', 'dataMax']} />
                    <Tooltip />
                    <Legend verticalAlign="top" height={36} />
                    <Line type="monotone" dataKey="max" stroke="#641ae6" strokeWidth={3} />
                </LineChart>
            </ResponsiveContainer>

            <div className="flex flex-wrap justify-center mt-8">
                <div className="stats shadow w-56 m-2 bg-base-200">
                    <div className="stat">
                        <div className="stat-title">Current Max</div>
                        <div className="stat-value">{stats.currentMax.max.toFixed(2)} kg</div>
                        <div className="stat-desc">{new Date(stats.currentMax.date).toLocaleDateString()}</div>
                    </div>
                </div>
                <div className="stats shadow w-56 m-2 bg-base-200">
                    <div className="stat">
                        <div className="stat-title">All Time Max</div>
                        <div className="stat-value">{stats.allTimeMax.max.toFixed(2)} kg</div>
                        <div className="stat-desc">{new Date(stats.allTimeMax.date).toLocaleDateString()}</div>
                    </div>
                </div>
            </div>
        </div>
    )
}
