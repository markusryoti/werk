import { Movement } from "../pages/workouts"

interface Props {
    movement: Movement
}

export default function RemoveMovement({ movement }: Props) {
    const deleteMovement = (id: number) => {
        console.log(id)
    }

    return (
        <div>
            <button onClick={() => deleteMovement(movement.id)} className="btn btn-error w-full">
                Remove movement
            </button>
        </div>
    )
}
