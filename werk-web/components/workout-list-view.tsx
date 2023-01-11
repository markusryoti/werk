import Link from "next/link"
import { IWorkout } from "../pages/workouts"

interface Props {
    workout: IWorkout
}

export default function WorkoutListView({ workout }: Props) {
    return (
        <div className="card bg-base-200 mt-2">
            <div className="card-body">
                <Link href={`/workouts/${workout.id}`}>
                    <h2 className="card-title">{workout.date.toLocaleString()}</h2>
                </Link>
                <p>{workout.name}</p>
            </div>
        </div>
    )
}
