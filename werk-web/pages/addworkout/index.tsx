import { useRouter } from "next/router"
import { FormEvent, useState } from "react"
import { useClientRequest } from "../../hooks/use-request"

export default function AddWorkout() {
    const [workoutName, setWorkoutName] = useState('')

    const router = useRouter()
    const { doRequest } = useClientRequest()

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()

        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts`

        doRequest(url, 'POST', { workoutName: workoutName })
            .then(res => res.json())
            .then(res => console.log(res))
            .then(() => router.push('workouts'))
            .catch(err => console.error(err))
    }

    return (
        <div className="flex justify-center p-1">
            <div className="card w-full md:w-1/2 bg-base-200 shadow-xl mt-16">
                <h1 className="card-title p-2">Add workout</h1>
                <form onSubmit={handleSubmit} className="">
                    <div className="flex flex-col m-2">
                        <label htmlFor="workoutName">Workout Name</label>
                        <input type="text" id="workoutName" className="input" onChange={e => setWorkoutName(e.target.value.trim())} />
                    </div>
                    <div className="flex justify-center p-2">
                        <button type="submit" className="btn btn-primary">Add</button>
                    </div>
                </form>
            </div>
        </div>
    )
}
