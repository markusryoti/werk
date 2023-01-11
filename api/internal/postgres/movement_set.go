package postgres

import "github.com/markusryoti/werk/internal/types"

func (p *PostgresRepo) AddMovementSet(movementId uint64, set types.Set) error {
	_, err := p.db.NamedExec(`INSERT INTO movement_set (reps, weight, movement_id) VALUES (:reps, :weight, :movementId)`,
		map[string]interface{}{
			"reps":       set.Reps,
			"weight":     set.Weight,
			"movementId": movementId,
		})

	if err != nil {
		return err
	}

	return nil
}
