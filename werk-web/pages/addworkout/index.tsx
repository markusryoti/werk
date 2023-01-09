import { FormEvent, useState } from "react"
import { useAuth } from "../context/AuthUserContext"

export default function AddWorkout() {
    const [workoutName, setWorkoutName] = useState('')

    const { getToken } = useAuth()

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()

        const token = await getToken()

        fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/add`, {
            method: 'POST',
            body: JSON.stringify({
                "workoutName": workoutName
            }),
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
        })
            .then(res => res.json())
            .then(res => console.log(res))
            .catch(err => console.error(err))
    }

    return (
        <div>
            <h1>Add workout</h1>
            <form onSubmit={handleSubmit}>
                <label htmlFor="workoutName">Workout Name</label>
                <input type="text" id="workoutName" onChange={e => setWorkoutName(e.target.value.trim())} />
                <button type="submit">Submit</button>
            </form>
        </div>
    )
}
