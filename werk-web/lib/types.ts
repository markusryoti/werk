export interface IWorkout {
    id: number;
    date: Date;
    name: string;
    movements: Movement[]
}

export interface Movement {
    id: number;
    name: string;
    sets: ISet[]
}

export interface ISet {
    id: number;
    reps: number;
    weight: number;
}
