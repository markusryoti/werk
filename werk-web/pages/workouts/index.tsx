import { useEffect, useState } from "react"
import { useAuth } from "../context/AuthUserContext"
import Workout from "./workout";

export interface IWorkout {
    id: number;
    date: Date;
    name: string;
    movements: Movement[]
}

export interface Movement {

}

export default function Workouts() {
    const [workouts, setWorkouts] = useState([])
    const { getToken } = useAuth()

    useEffect(() => {
        getToken().then(token => {
            fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/workouts`, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
            })
                .then(res => res.json())
                .then(wouts => setWorkouts(wouts))
                .catch(err => console.error(err))
        })
            .catch(err => console.error(err))
    }, [])

    return (
        <>
            <div>Workouts</div>
            {workouts && workouts.map((w, i) => <Workout key={i} workout={w} />)}
        </>
    )
}

