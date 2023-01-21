package postgres

import "github.com/markusryoti/werk/internal/types"

func (p *PostgresRepo) AddMovementSet(movementId uint64, set types.Set) error {
	_, err := p.db.NamedExec(`INSERT INTO movement_set (reps, weight, movement_id, user_id) VALUES (:reps, :weight, :movementId, :userId)`,
		map[string]interface{}{
			"reps":       set.Reps,
			"weight":     set.Weight,
			"movementId": movementId,
			"userId":     set.User,
		})

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepo) GetMovementSet(movementSetId uint64) (types.Set, error) {
	var set types.Set

	row := p.db.QueryRow("SELECT id, user_id FROM movement_set WHERE id = $1", movementSetId)
	err := row.Scan(&set.ID, &set.User)

	return set, err
}

func (p *PostgresRepo) DeleteMovementSet(movementSetId uint64) error {
	_, err := p.db.NamedExec(`DELETE FROM movement_set WHERE id = :movementSetId`,
		map[string]interface{}{
			"movementSetId": movementSetId,
		})

	return err
}
