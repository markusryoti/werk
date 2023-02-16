import { FormEvent, useState } from "react"
import { useClientRequest } from "../hooks/use-request";
import { ISet } from "../lib/types";

interface Props {
    workoutId: number;
    movementId: number;
    updateWorkout: () => void
    addSetToMovement: (movementId: number, set: ISet) => void
}

export default function AddSet({ workoutId, movementId, addSetToMovement }: Props) {
    const [reps, setReps] = useState('')
    const [weight, setWeight] = useState('')

    const { doRequest } = useClientRequest()

    const addSet = async (movementId: number, e: FormEvent) => {
        e.preventDefault()

        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${workoutId}/workoutMovements/${movementId}`
        const res = await doRequest(url, 'POST', { reps: parseInt(reps), weight: parseInt(weight) })
        const newSet = await res.json()

        if (res.ok) {
            addSetToMovement(movementId, newSet)
            setReps('')
            setWeight('')
        }
    }

    return (
        <form onSubmit={(e) => addSet(movementId, e)} className="flex flex-col md:flex-row p-2">
            <div className="flex flex-col p-1">
                <label htmlFor="reps">Reps</label>
                <input value={reps} onChange={e => setReps(e.target.value)} id="reps" type="number" className="input" />
            </div>
            <div className="flex flex-col p-1">
                <label htmlFor="weight">Weight</label>
                <input value={weight} onChange={e => setWeight(e.target.value)} id="weight" type="number" className="input" />
            </div>
            <div className="flex items-end p-1">
                <button className="btn btn-secondary mt-4 w-full">
                    Add Set
                </button>
            </div>
        </form>
    )
}
