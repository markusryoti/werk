import { FormEvent, useState } from "react";
import { useClientRequest } from "../hooks/use-request"

interface Props {
    workoutId: number;
    getWorkout: () => void;
}

export default function AddMovement({ workoutId, getWorkout }: Props) {
    const [movementName, setMovementName] = useState('')

    const { doRequest } = useClientRequest()

    const addMovement = async (e: FormEvent) => {
        e.preventDefault()

        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${workoutId}/addMovement`
        const res = await doRequest(url, 'POST', { movementName: movementName })

        if (res.ok) {
            getWorkout()
        }

        setMovementName('')
    }

    return (
        <div className="card border border-base-300 bg-base-200 shadow-l p-4 mb-4">
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
    )
}
