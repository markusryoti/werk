import { FormEvent, useState } from "react"
import { useAuth } from "../pages/context/AuthUserContext"

interface Props {
    workoutId: number;
    movementId: number;
    updateWorkout: (token: string) => void
}

export default function AddSet({ workoutId, movementId, updateWorkout }: Props) {
    const [reps, setReps] = useState(0)
    const [weight, setWeight] = useState(0)

    const { getToken } = useAuth()

    const addSet = async (movementId: number, e: FormEvent) => {
        e.preventDefault()

        const token = await getToken()

        const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${workoutId}/workoutMovements/${movementId}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            }, body: JSON.stringify({
                reps: reps,
                weight: weight
            })
        })

        if (res.ok) {
            updateWorkout(token!)
            setReps(0)
            setWeight(0)
        }
    }

    return (
        <form onSubmit={(e) => addSet(movementId, e)} className="flex flex-col md:flex-row p-2 mt-2">
            <div className="flex flex-col">
                <label htmlFor="reps">Reps</label>
                <input value={reps} onChange={e => setReps(parseInt(e.target.value))} id="reps" type="number" className="input" />
            </div>
            <div className="flex flex-col">
                <label htmlFor="weight">Weight</label>
                <input value={weight} onChange={e => setWeight(parseInt(e.target.value))} id="weight" type="number" className="input" />
            </div>
            <button className="btn btn-secondary mt-4">
                Add Set
            </button>
        </form>
    )
}
