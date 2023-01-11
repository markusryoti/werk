import { useRouter } from "next/router"
import { FormEvent, useState } from "react"
import { useAuth } from "../context/AuthUserContext"

export default function AddWorkout() {
    const [workoutName, setWorkoutName] = useState('')

    const router = useRouter()
    const { getToken } = useAuth()

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()

        const token = await getToken()

        fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts`, {
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
            .then(() => router.push('workouts'))
            .catch(err => console.error(err))
    }

    return (
        <div className="flex justify-center">
            <div className="card w-1/2 bg-base-200 shadow-xl mt-16">
                <div className="p-8">
                    <h1 className="text-2xl">Add workout</h1>
                </div>
                <form onSubmit={handleSubmit} className="p-8">
                    <div className="flex items-center m-2">
                        <label htmlFor="workoutName" className="flex-none pr-2">Workout Name</label>
                        <input type="text" id="workoutName" className="input w-full" onChange={e => setWorkoutName(e.target.value.trim())} />
                    </div>
                    <div className="flex justify-center">
                        <div className="m-2">
                            <button type="submit" className="btn btn-primary">Add</button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    )
}
