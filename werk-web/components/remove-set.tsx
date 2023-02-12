import { useClientRequest } from "../hooks/use-request"
import { Set } from "../lib/types"

interface Props {
    set: Set
    updateWorkout: () => void;
}

export default function RemoveSet({ set, updateWorkout }: Props) {
    const { doRequest } = useClientRequest()

    const deleteSet = async () => {
        const url = `${process.env.NEXT_PUBLIC_BACKEND_URL}/sets/${set.id}`
        const res = await doRequest(url, 'DELETE')

        if (res.ok) {
            updateWorkout()
        }
    }

    return (
        <div onClick={deleteSet} className="text-error">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor" className="w-4 h-4">
                <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
        </div>
    )
}
