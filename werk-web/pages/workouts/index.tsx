import { GetServerSidePropsContext } from "next";
import WorkoutListView from "../../components/workout-list-view";
import { IWorkout } from "../../lib/types";

interface Props {
    workouts: IWorkout[]
}

export default function Workouts({ workouts }: Props) {
    return (
        <div className="container flex flex-col items-center pt-8 p-1 m-auto">
            {workouts.length > 0
                ? workouts.map((w: IWorkout, i: number) => <WorkoutListView key={i} workout={w} />)
                : <p>No workouts</p>}
        </div>
    )
}

export async function getServerSideProps(context: GetServerSidePropsContext) {
    const token = context.req.cookies.session

    const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        }
    })

    if (res.status === 401) {
        return {
            redirect: {
                destination: '/login'
            }
        }
    }

    const workouts = await res.json()

    return {
        props: {
            workouts: workouts ? workouts : []
        }
    }
}

