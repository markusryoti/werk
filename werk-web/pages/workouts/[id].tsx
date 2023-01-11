import { GetServerSidePropsContext } from "next"
import { IWorkout } from "."

interface Props {
    workout: IWorkout
}

export default function WorkoutDetail({ workout }: Props) {
    return (
        <div>{JSON.stringify(workout)}</div>
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
