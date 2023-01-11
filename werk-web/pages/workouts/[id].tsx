import { GetServerSidePropsContext } from "next"
import { IWorkout } from "."

interface Props {
    workout: IWorkout
}

export default function WorkoutDetail({ workout }: Props) {
    return (
        <div className="container mx-auto mt-8">
            <div className="mt-4 mb-4 p-2">
                <h2 className="text-2xl font-bold">{workout.date.toLocaleString()}</h2>
                <h3 className="text-xl">{workout.name}</h3>
            </div>
            {workout.movements && workout.movements.map((movement, i) => {
                return (
                    <div key={`movement-${i}`} className="collapse collapse-arrow border border-base-300 bg-base-200 rounded-box mt-4 mb-4">
                        <input type="checkbox" className="peer" />
                        <div className="collapse-title">
                            {movement.name}
                        </div>
                        <div className="collapse-content">
                            <table className="table w-full">
                                <thead>
                                    <tr>
                                        <th>Set</th>
                                        <th>Reps</th>
                                        <th>Weight</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {movement.sets.map((set, j) => {
                                        return (
                                            <tr key={`set-${j}`}>
                                                <td>{j + 1}</td>
                                                <td>{set.reps}</td>
                                                <td>{set.weight}</td>
                                            </tr>
                                        )
                                    })}
                                </tbody>
                            </table>
                        </div>
                    </div >
                )
            })}
        </div >
    )
}

export async function getServerSideProps(context: GetServerSidePropsContext) {
    const id = context.params?.id
    const token = context.req.cookies.session


    const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts/${id}`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        }
    })

    const workout = await res.json()

    return {
        props: {
            workout
        }
    }
}
