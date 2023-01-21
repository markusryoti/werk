import { useRouter } from "next/router"
import { FormEvent, useEffect, useState } from "react"
import { IWorkout } from "."
import AddSet from "../../components/add-set"
import RemoveMovement from "../../components/remove-movement"
import RemoveSet from "../../components/remove-set"
import Spinner from "../../components/spinner"
import { useClientRequest } from "../../hooks/use-request"
import { parseDate } from "../../utils/date"

export default function WorkoutDetail() {
    const [workout, setWorkout] = useState<IWorkout>()
    const [movementName, setMovementName] = useState('')

    const { doRequest } = useClientRequest()

    const router = useRouter()

    const { id } = router.query

    useEffect(() => {
        getWorkout()
    }, [])

    const getWorkout = () => {
        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${id}`;

        doRequest(url, 'GET', undefined)
            .then(res => res.json())
            .then(w => setWorkout(w))
            .catch(err => console.error(err))
    }

    const addMovement = async (e: FormEvent) => {
        e.preventDefault()

        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${id}/addMovement`
        const res = await doRequest(url, 'POST', { movementName: movementName })

        if (res.ok) {
            getWorkout()
        }

        setMovementName('')
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
            {workout && workout.movements.map((movement, i) => {
                return (
                    <div key={`movement-${i}`} className="collapse collapse-arrow border border-base-300 bg-base-200 rounded-box mt-4 mb-4">
                        <input type="checkbox" className="peer" />
                        <div className="collapse-title">
                            {movement.name}
                        </div>
                        <div className="collapse-content flex flex-col align-center">
                            {movement.sets.length > 0 ? (
                                <table className="table w-full">
                                    <thead>
                                        <tr>
                                            <th>Set</th>
                                            <th>Reps</th>
                                            <th>Weight</th>
                                            <th>Delete</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        {movement.sets.map((set, j) => {
                                            return (
                                                <tr key={`set-${j}`}>
                                                    <td>{j + 1}</td>
                                                    <td>{set.reps}</td>
                                                    <td>{set.weight}</td>
                                                    <td>
                                                        <RemoveSet set={set} />
                                                    </td>
                                                </tr>
                                            )
                                        })}
                                    </tbody>
                                </table>
                            ) : <p>No sets</p>}
                            <div className="divider"></div>
                            <AddSet workoutId={workout.id} movementId={movement.id} updateWorkout={getWorkout} />
                            <div className="divider"></div>
                            <RemoveMovement movement={movement} />
                        </div>
                    </div>
                )
            })}
            <div className="card bg-base-200 shadow-l p-4 mb-4">
                <h3 className="card-title">Add movement</h3>
                <form onSubmit={addMovement} className="flex flex-col">
                    <div className="p-2">
                        <input placeholder="Movement name" onChange={e => setMovementName(e.target.value)} value={movementName} className="input input-bordered w-full" />
                    </div>
                    <div className="p-2">
                        <button onClick={addMovement} className="btn btn-primary w-full">Add</button>
                    </div>
                </form>
            </div>
            <div className="card bg-base-200 shadow-l p-4">
                <h3 className="card-title mb-2">Remove Workout</h3>
                <div>
                    <button className="btn btn-error w-full">Remove</button>
                </div>
            </div>
        </div >
    )
}
