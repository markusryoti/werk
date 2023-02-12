import Link from "next/link"
import { IWorkout } from "../lib/types"
import { parseDate } from "../utils/date"

interface Props {
    workout: IWorkout
}

export default function WorkoutListView({ workout }: Props) {
    return (
        <div className="card border border-base-300 bg-base-200 mt-2 w-full md:w-1/2">
            <div className="card-body p-6">
                <Link href={`/workouts/${workout.id}`}>
                    <h2 className="card-title">{parseDate(workout.date)}</h2>
                </Link>
                <p>{workout.name}</p>
            </div>
        </div>
    )
}
