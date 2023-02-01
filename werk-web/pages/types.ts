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
    id: number;
    reps: number;
    weight: number;
}
