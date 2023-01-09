package postgres

import "github.com/markusryoti/werk/internal/types"

func (p *PostgresRepo) AddNewMovement(workoutId uint64, name string) error {
	_, err := p.db.NamedExec(`INSERT INTO movement (name, workout_id)
                                VALUES (:name, :workoutId)`,
		map[string]interface{}{
			"name":      name,
			"workoutId": workoutId,
		})

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepo) GetMovementsFromWorkout(workoutId uint64) ([]types.Movement, error) {
	query := `SELECT movement.id, name, workout_id, movement_set.id, reps, weight, movement_id
                FROM movement
                JOIN movement_set
                    ON movement.id = movement_set.movement_id
                WHERE movement.workout_id = $1
                ORDER BY movement_set.id ASC`

	rows, err := p.db.Query(query, workoutId)

	movements := make(map[uint64]Movement)
	sets := make([]MovementSet, 0)

	for rows.Next() {
		var (
			m Movement
			s MovementSet
		)

		err = rows.Scan(&m.ID, &m.Name, &m.WorkoutId,
			&s.ID, &s.Reps, &s.Weight, &s.MovementId)

		sets = append(sets, s)
		movements[m.ID] = m
	}

	err = rows.Err()

	return movementsToDomain(sets, movements), err
}

func movementsToDomain(sets []MovementSet, movementMap map[uint64]Movement) []types.Movement {
	var (
		movements = make([]types.Movement, 0)
	)

	var currentMovement Movement

	for _, set := range sets {
		if currentMovement.ID == 0 {
			currentMovement = movementMap[set.MovementId]
		}

		if currentMovement.ID != set.MovementId {
			movements = append(movements, currentMovement.toDomain())
			currentMovement = movementMap[set.MovementId]
		}

		currentMovement.Sets = append(currentMovement.Sets, set)
	}

	movements = append(movements, currentMovement.toDomain())

	return movements
}
