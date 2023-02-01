package postgres

import (
	"sort"

	"github.com/markusryoti/werk/internal/types"
)

func (p *PostgresRepo) AddNewWorkout(userId, name string) error {
	_, err := p.db.NamedExec(`INSERT INTO workout (user_id, name) VALUES (:userId, :name)`,
		map[string]interface{}{
			"userId": userId,
			"name":   name,
		})

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepo) GetWorkoutsByUser(userId string) ([]types.Workout, error) {
	rows, err := p.db.Queryx(`SELECT id, created_at, name, user_id FROM workout
                                WHERE user_id = $1
                                ORDER BY created_at DESC`, userId)

	workouts := make([]Workout, 0)

	for rows.Next() {
		var w Workout
		err = rows.StructScan(&w)
		workouts = append(workouts, w)
	}

	err = rows.Err()

	return p.workoutsToDomain(workouts), err
}

func (p *PostgresRepo) workoutsToDomain(workouts []Workout) []types.Workout {
	domainWorkouts := make([]types.Workout, len(workouts))

	for i, w := range workouts {
		domainWorkouts[i] = w.toDomain()
	}

	return domainWorkouts
}

func (p *PostgresRepo) GetWorkout(workoutId uint64) (types.Workout, error) {
	query := `
        SELECT
          workout.created_at,
          workout.name AS workout_name,
          workout.user_id,
          workout_movement.id AS workout_movement_id,
          movement.name AS movement_name,
          movement_set.id AS movement_set_id,
          movement_set.reps,
          movement_set.weight
        FROM
          workout
          JOIN workout_movement ON workout.id = workout_movement.workout_id
          JOIN movement ON workout_movement.movement_id = movement.id
          LEFT JOIN movement_set ON workout_movement.id = movement_set.workout_movement_id
        WHERE
          workout.id = $1
        ORDER BY
          workout_movement.id ASC`

	rows, err := p.db.Queryx(query, workoutId)
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
			movementName  string
			movementId    uint64
			movementSetId uint64
			reps          uint8
			weight        int
		)

		err = rows.Scan(&workout.Date, &workout.Name, &workout.User, &movementId, &movementName, &movementSetId, &reps, &weight)

		if movementName != currentMovement.Name {
			if currentMovement.Name != "" {
				currentMovement.Sets = p.sortSets(currentMovement.Sets)
				workout.Movements = append(workout.Movements, currentMovement)
			}

			currentMovement = types.Movement{ID: movementId, Name: movementName}
			currentMovement.Sets = make([]types.Set, 0)
		}

		if reps != 0 && weight != 0 {
			currentMovement.Sets = append(currentMovement.Sets, types.Set{ID: movementSetId, Reps: reps, Weight: weight})
		}
	}

	if currentMovement.ID != 0 {
		currentMovement.Sets = p.sortSets(currentMovement.Sets)
		workout.Movements = append(workout.Movements, currentMovement)
	}

	workout.ID = workoutId

	return workout, rows.Err()
}

func (p *PostgresRepo) sortSets(sets []types.Set) []types.Set {
	sorted := sets

	sort.Slice(sets, func(i, j int) bool {
		return sets[i].ID < sets[j].ID
	})

	return sorted
}

func (p *PostgresRepo) DeleteWorkout(workoutId uint64) error {
	_, err := p.db.NamedExec(`DELETE FROM workout WHERE id = :workoutId`,
		map[string]interface{}{
			"workoutId": workoutId,
		})

	return err
}
