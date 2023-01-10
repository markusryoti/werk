import Workout from "./workout";

export interface IWorkout {
    id: number;
    date: Date;
    name: string;
    movements: Movement[]
}

export interface Movement {

}

export default function Workouts({ workouts }: any) {
    return (
        <>
            <div>Workouts</div>
            {workouts && workouts.map((w, i) => <Workout key={i} workout={w} />)}
        </>
    )
}

export async function getServerSideProps(context: any) {
    const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts`, { credentials: "include" })
    const workouts = await res.json()

    console.log('cookies', context.req.cookies)

    console.log(workouts)

    return {
        props: {
            workouts: []
        }
    }
}

