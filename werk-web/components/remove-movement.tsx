import { useClientRequest } from "../hooks/use-request";
import { Movement } from "../lib/types";

interface Props {
    movement: Movement;
    updateWorkout: () => void;
}

export default function RemoveMovement({ movement, updateWorkout }: Props) {
    const { doRequest } = useClientRequest()

    const deleteMovement = async () => {
        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/movements/${movement.id}`
        const res = await doRequest(url, 'DELETE')

        if (res.ok) {
            updateWorkout()
        }
    }

    return (
        <div>
            <button onClick={deleteMovement} className="btn btn-error w-full">
                Remove movement
            </button>
        </div>
    )
}
