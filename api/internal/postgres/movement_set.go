package postgres

import "github.com/markusryoti/werk/internal/types"

func (p *PostgresRepo) AddMovementSet(workoutMovementId uint64, set types.Set) (types.Set, error) {
	var (
		id     uint64
		newSet types.Set
	)

	err := p.db.QueryRowx(`INSERT INTO movement_set (reps, weight, workout_movement_id, user_id) 
                            VALUES ($1, $2, $3, $4)
                            RETURNING id, reps, weight, user_id`,
		set.Reps, set.Weight, workoutMovementId, set.User).
		Scan(&id, &newSet.Reps, &newSet.Weight, &newSet.User)

	if err != nil {
		return types.Set{}, err
	}

	// Why do I need to set it like this?
	set.ID = id

	return set, nil
}

func (p *PostgresRepo) GetMovementSet(movementSetId uint64) (types.Set, error) {
	var set types.Set

	row := p.db.QueryRowx("SELECT id, reps, weight, user_id FROM movement_set WHERE id = $1", movementSetId)
	err := row.Scan(&set.ID, &set.Reps, &set.Weight, &set.User)

	return set, err
}

func (p *PostgresRepo) DeleteMovementSet(movementSetId uint64) error {
	_, err := p.db.NamedExec(`DELETE FROM movement_set WHERE id = :movementSetId`,
		map[string]interface{}{
			"movementSetId": movementSetId,
		})

	return err
}
