import { IWorkout } from ".";

interface Props {
    workout: IWorkout
}

export default function Workout({ workout }: Props) {
    return (
        <div>
            {workout.name}
        </div>
    )
}
