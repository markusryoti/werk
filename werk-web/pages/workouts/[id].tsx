import { useRouter } from "next/router"
import { FormEvent, useEffect, useState } from "react"
import { IWorkout } from "."
import AddSet from "../../components/add-set"
import Spinner from "../../components/spinner"
import { useAuth } from "../context/AuthUserContext"

export default function WorkoutDetail() {
    const [workout, setWorkout] = useState<IWorkout>()
    const [movementName, setMovementName] = useState('')

    const router = useRouter()
    const { getToken } = useAuth()

    const { id } = router.query

    useEffect(() => {
        getToken().then(token => {
            token && getWorkout(token)
        }).catch(err => console.error(err))
    }, [])

    const getWorkout = (token: string) => {
        fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${id}`, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            }
        }).then(res => res.json())
            .then(w => setWorkout(w))
            .catch(err => console.error(err))
    }

    const addMovement = async (e: FormEvent) => {
        e.preventDefault()

        const token = await getToken()

        const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${id}/addMovement`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            }, body: JSON.stringify({
                movementName: movementName
            })
        })

        if (res.ok && token) {
            getWorkout(token)
        }

        setMovementName('')
    }


    if (!workout) {
        return <Spinner />
    }

    console.log(workout)

    return (
        <div className="container mx-auto mt-8 p-1">
            <div className="mt-4 mb-4 p-2">
                <h2 className="text-2xl font-bold">{workout.date.toLocaleString()}</h2>
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
                            <table className="table w-full">
                                <thead>
                                    <tr>
                                        <th>Set</th>
                                        <th>Reps</th>
                                        <th>Weight</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {movement.sets.map((set, j) => {
                                        return (
                                            <tr key={`set-${j}`}>
                                                <td>{j + 1}</td>
                                                <td>{set.reps}</td>
                                                <td>{set.weight}</td>
                                            </tr>
                                        )
                                    })}
                                </tbody>
                            </table>
                            <AddSet workoutId={workout.id} movementId={movement.id} updateWorkout={getWorkout} />
                        </div>
                    </div>
                )
            })}
            <div className="card bg-base-200 shadow-xl p-4">
                <h3 className="card-title">Add movement</h3>
                <form onSubmit={addMovement} className="flex flex-col">
                    <div className="p-2">
                        <input placeholder="Add new movement" onChange={e => setMovementName(e.target.value)} value={movementName} className="input input-bordered w-full" />
                    </div>
                    <div className="p-2">
                        <button onClick={addMovement} className="btn btn-primary">Add</button>
                    </div>
                </form>
            </div>
        </div >
    )
}
