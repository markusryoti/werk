import { Movement } from "../pages/workouts"
import AddSet from "./add-set"
import RemoveMovement from "./remove-movement"
import RemoveSet from "./remove-set"

interface Props {
    movement: Movement
    workoutId: number
    getWorkout: () => void
    key: string
}

export default function MovementCollapse({ movement, workoutId, getWorkout, key }: Props) {
    return (
        <div key={key} className="collapse collapse-arrow border border-base-300 bg-base-200 rounded-box mt-4 mb-4">
            <input type="checkbox" className="peer" />
            <div className="collapse-title">
                {movement.name}
            </div>
            <div className="collapse-content flex flex-col align-center">
                {movement.sets.length > 0 ? (
                    <table className="table w-full">
                        <thead>
                            <tr>
                                <th>Set</th>
                                <th>Reps</th>
                                <th>Weight</th>
                                <th>Delete</th>
                            </tr>
                        </thead>
                        <tbody>
                            {movement.sets.map((set, j) => {
                                return (
                                    <tr key={`set-${j}`}>
                                        <td>{j + 1}</td>
                                        <td>{set.reps}</td>
                                        <td>{set.weight}</td>
                                        <td>
                                            <RemoveSet set={set} updateWorkout={getWorkout} />
                                        </td>
                                    </tr>
                                )
                            })}
                        </tbody>
                    </table>
                ) : <p>No sets</p>}
                <div className="divider"></div>
                <AddSet workoutId={workoutId} movementId={movement.id} updateWorkout={getWorkout} />
                <div className="divider"></div>
                <RemoveMovement movement={movement} updateWorkout={getWorkout} />
            </div>
        </div>
    )
}
