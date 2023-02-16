import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import AddMovement from "../../components/add-movement"
import MovementCollapse from "../../components/movement-collapse"
import RemoveWorkout from "../../components/remove-workout"
import Spinner from "../../components/spinner"
import { useAuth } from "../../context/AuthUserContext"
import { useClientRequest } from "../../hooks/use-request"
import { ISet, IWorkout, Movement } from "../../lib/types"
import { parseDate } from "../../utils/date"

export default function WorkoutDetail() {
    const [workout, setWorkout] = useState<IWorkout>()

    const router = useRouter()
    const { id } = router.query
    const { doRequest } = useClientRequest()

    const { authUser } = useAuth()

    useEffect(() => {
        if (authUser) {
            getWorkout()
        } else {
            router.push("/login")
        }
        // eslint-disable-next-line
    }, [])

    const getWorkout = () => {
        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${id}`;

        doRequest(url, 'GET')
            .then(res => res.json())
            .then(w => setWorkout(w))
            .catch(err => console.error(err))
    }

    const addSetToMovement = (movementId: number, newSet: ISet) => {
        if (!workout) return

        const updatedMovements = workout.movements.map(movement => {
            if (movement.id === movementId) {
                movement.sets.push(newSet)
            }

            return movement
        })

        setWorkout({
            ...workout,
            movements: updatedMovements
        })
    }

    const addMovement = (newMovement: Movement) => {
        if (!workout) return

        const updatedMovements = [...workout.movements, newMovement]

        setWorkout({
            ...workout,
            movements: updatedMovements
        })
    }

    if (!workout) {
        return <Spinner />
    }

    return (
        <div className="flex justify-center container mx-auto mt-8 p-1">
            <div className="flex flex-col w-full md:w-1/2">
                <div className="mt-4 mb-4 p-2">
                    <h2 className="text-2xl font-bold">{parseDate(workout.date)}</h2>
                    <h3 className="text-xl">{workout.name}</h3>
                </div>
                {workout && workout.movements.map((movement) => {
                    return <MovementCollapse
                        movement={movement}
                        workoutId={workout.id}
                        getWorkout={getWorkout}
                        addSetToMovement={addSetToMovement}
                        key={`${movement.id}`}
                    />
                })}
                <AddMovement workoutId={workout.id} addMovement={addMovement} />
                <RemoveWorkout workoutId={workout.id} />
            </div>
        </div >
    )
}
