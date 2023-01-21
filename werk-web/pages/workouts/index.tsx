import WorkoutListView from "../../components/workout-list-view";

export interface IWorkout {
    id: number;
    date: Date;
    name: string;
    movements: Movement[]
}

export interface Movement {
    id: number;
    name: string;
    sets: Set[]
}

export interface Set {
    reps: number;
    weight: number;
}

export default function Workouts({ workouts }: any) {
    return (
        <div className="container mx-auto pt-8 p-1">
            {workouts && workouts.map((w: IWorkout, i: number) => <WorkoutListView key={i} workout={w} />)}
        </div>
    )
}

export async function getServerSideProps(context: any) {
    const token = context.req.cookies.session

    const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        }
    })

    const workouts = await res.json()

    return {
        props: {
            workouts: workouts ? workouts : []
        }
    }
}

