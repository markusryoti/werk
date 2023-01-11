package postgres

import "github.com/markusryoti/werk/internal/types"

func (p *PostgresRepo) AddNewWorkout(name string) error {
	_, err := p.db.NamedExec(`INSERT INTO workout (name) VALUES (:name)`,
		map[string]interface{}{
			"name": name,
		})

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepo) GetAllWorkouts() ([]types.Workout, error) {
	rows, err := p.db.Queryx(`SELECT * FROM workout
                                ORDER BY date DESC`)

	workouts := make([]Workout, 0)

	for rows.Next() {
		w := Workout{}
		err = rows.StructScan(&w)
		workouts = append(workouts, w)
	}

	err = rows.Err()

	return workoutsToDomain(workouts), err
}

func workoutsToDomain(workouts []Workout) []types.Workout {
	domainWorkouts := make([]types.Workout, 0)

	for _, w := range workouts {
		domainWorkouts = append(domainWorkouts, w.toDomain())
	}

	return domainWorkouts
}

func (p *PostgresRepo) GetWorkout(workoutId uint64) (types.Workout, error) {
	query := `SELECT date, workout.name, movement.id, movement.name, reps, weight
                FROM workout
                JOIN movement
                    ON workout.id = movement.workout_id
                JOIN movement_set
                    ON movement.id = movement_set.movement_id
                WHERE workout.id = $1`

	rows, err := p.db.Query(query, workoutId)
	if err != nil {
		return types.Workout{}, err
	}

	var (
		workout         types.Workout
		currentMovement types.Movement
	)

	workout.Movements = make([]types.Movement, 0)
	currentMovement.Sets = make([]types.Set, 0)

	for rows.Next() {
		var (
			movementName string
			movementId   uint64
			reps         uint8
			weight       uint8
		)

		err = rows.Scan(&workout.Date, &workout.Name, &movementId, &movementName, &reps, &weight)

		if movementName != currentMovement.Name {
			if currentMovement.Name != "" {
				workout.Movements = append(workout.Movements, currentMovement)
			}

			currentMovement = types.Movement{ID: movementId, Name: movementName}
			currentMovement.Sets = make([]types.Set, 0)
		}

		currentMovement.Sets = append(currentMovement.Sets, types.Set{Reps: reps, Weight: weight})
	}

	workout.Movements = append(workout.Movements, currentMovement)
	workout.ID = workoutId

	return workout, rows.Err()
}
