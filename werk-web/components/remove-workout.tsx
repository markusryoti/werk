import { useRouter } from "next/router"
import { useClientRequest } from "../hooks/use-request"

interface Props {
    workoutId: number;
}

export default function RemoveWorkout({ workoutId }: Props) {
    const { doRequest } = useClientRequest()

    const router = useRouter()

    const removeWorkout = async () => {
        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${workoutId}`
        const res = await doRequest(url, 'DELETE')

        if (res.ok) {
            router.push("/workouts")
        }
    }

    return (
        <div className="card border border-base-300 bg-base-200 shadow-l p-4">
            <h3 className="card-title text-xl font-medium mb-2">Remove Workout</h3>
            <div>
                <button onClick={removeWorkout} className="btn btn-error w-full">Remove</button>
            </div>
        </div>
    )
}
