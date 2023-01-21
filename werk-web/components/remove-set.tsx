import { Set } from "../pages/workouts"

interface Props {
    set: Set
}

export default function RemoveSet({ set }: Props) {
    const deleteSet = (id: number) => {
        console.log(id)
    }

    return (
        <div onClick={() => deleteSet(set.id)} className="text-error">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor" className="w-4 h-4">
                <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
        </div>
    )
}
