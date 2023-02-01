import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import AddMovement from "../../components/add-movement"
import MovementCollapse from "../../components/movement-collapse"
import RemoveWorkout from "../../components/remove-workout"
import Spinner from "../../components/spinner"
import { useClientRequest } from "../../hooks/use-request"
import { parseDate } from "../../utils/date"
import { IWorkout } from "../types"

export default function WorkoutDetail() {
    const [workout, setWorkout] = useState<IWorkout>()

    const router = useRouter()
    const { doRequest } = useClientRequest()

    const { id } = router.query

    useEffect(() => {
        getWorkout()
    }, [])

    const getWorkout = () => {
        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${id}`;

        doRequest(url, 'GET')
            .then(res => res.json())
            .then(w => setWorkout(w))
            .catch(err => console.error(err))
    }

    if (!workout) {
        return <Spinner />
    }

    return (
        <div className="container mx-auto mt-8 p-1">
            <div className="mt-4 mb-4 p-2">
                <h2 className="text-2xl font-bold">{parseDate(workout.date)}</h2>
                <h3 className="text-xl">{workout.name}</h3>
            </div>
            {workout && workout.movements.map((movement) => {
                return <MovementCollapse
                    movement={movement}
                    workoutId={workout.id}
                    getWorkout={getWorkout}
                    key={`${movement.id}`}
                />
            })}
            <AddMovement workoutId={workout.id} getWorkout={getWorkout} />
            <RemoveWorkout workoutId={workout.id} />
        </div >
    )
}
