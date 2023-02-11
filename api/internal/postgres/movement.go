package postgres

import (
	"database/sql"
	"sort"
	"time"

	"github.com/markusryoti/werk/internal/maxestimate"
	"github.com/markusryoti/werk/internal/types"
)

func (p *PostgresRepo) AddNewMovement(workoutId uint64, newMovement types.Movement) error {
	var (
		movement Movement
		err      error
	)

	movement, err = p.getMovement(newMovement)
	if err != nil {
		if err == sql.ErrNoRows {
			movement, err = p.addMovement(newMovement)
			if err != nil {
				return err
			}
		} else {
			p.logger.Errorw("Unexpected error getting movement", "error", err)
			return err
		}
	}

	_, err = p.getMovementFromWorkout(movement.ID, workoutId)
	if err == nil {
		return nil
	}

	err = p.addMovementToWorkout(movement.ID, workoutId, newMovement.User)

	return err
}

func (p *PostgresRepo) getMovement(m types.Movement) (Movement, error) {
	query := `SELECT
                    movement.id,
                    movement.name,
                    movement.user_id
                FROM movement
                WHERE movement.name = $1 AND movement.user_id = $2`

	row := p.db.QueryRowx(query, m.Name, m.User)

	var movement Movement

	err := row.Scan(&movement.ID, &movement.Name, &movement.UserId)
	if err != nil && err != sql.ErrNoRows {
		p.logger.Info("Error getting movement", "movement", movement)
	}

	return movement, err
}

func (p *PostgresRepo) getMovementFromWorkout(movementId, workoutId uint64) (WorkoutMovement, error) {
	query := `SELECT
                    workout_movement.id,
                    workout_movement.workout_id,
                    workout_movement.movement_id,
                    workout_movement.user_id
                FROM workout_movement
                JOIN workout
                    ON workout_movement.workout_id = workout.id
                WHERE workout_movement.movement_id = $1 AND workout.id = $2`

	row := p.db.QueryRowx(query, movementId, workoutId)

	var movement WorkoutMovement

	err := row.Scan(&movement.ID, &movement.WorkoutId, &movement.MovementId, &movement.UserId)
	if err != nil && err != sql.ErrNoRows {
		p.logger.Errorw("Couldn't get movement from workout", "error", err,
			"movementId", movementId, "workoutId", workoutId)
	}

	return movement, err
}

func (p *PostgresRepo) addMovement(newMovement types.Movement) (Movement, error) {
	_, err := p.db.NamedExec(`INSERT INTO movement (name, user_id)
                                VALUES (:name, :userId)`,
		map[string]interface{}{
			"name":   newMovement.Name,
			"userId": newMovement.User,
		})

	if err != nil {
		p.logger.Errorw("Couldn't add new movement", "error", err, "movement", newMovement)
		return Movement{}, err
	}

	return p.getMovement(newMovement)
}

func (p *PostgresRepo) addMovementToWorkout(movementId, workoutId uint64, userId string) error {
	_, err := p.db.NamedExec(`INSERT INTO workout_movement (workout_id, movement_id, user_id)
                                VALUES (:workoutId, :movementId, :userId)`,
		map[string]interface{}{
			"workoutId":  workoutId,
			"movementId": movementId,
			"userId":     userId,
		})

	if err != nil {
		p.logger.Errorw("Couldn't add new workout movement", "error", err,
			"movementId", movementId, "workoutId", workoutId, "userId", userId)
	}

	return err
}

func (p *PostgresRepo) GetMovementsFromWorkout(workoutId uint64) ([]types.Movement, error) {
	query := `SELECT 
                    movement.id, 
                    movement.name, 
                    movement_set.id, 
                    movement_set.reps, 
                    movement_set.weight, 
                FROM movement
                JOIN workout_movement
                    ON movement.id = workout_movement.movement_id
                LEFT JOIN movement_set
                    ON workout_movement.id = movement_set.workout_movement_id
                WHERE workout_movement.workout_id = $1
                    ORDER BY movement_set.id ASC`

	rows, err := p.db.Queryx(query, workoutId)
	if err != nil {
		return nil, err
	}

	var (
		currentMovement types.Movement
		movements       = make([]types.Movement, 0)
	)

	for rows.Next() {
		var (
			set      types.Set
			movement types.Movement
		)

		err = rows.Scan(&movement.ID, &movement.Name, &set.ID, &set.Reps, &set.Weight)

		if movement.ID != currentMovement.ID {
			if currentMovement.ID != 0 {
				movements = append(movements, currentMovement)
				currentMovement = movement
			}
		}

		currentMovement.Sets = append(currentMovement.Sets, set)
	}

	if currentMovement.ID != 0 {
		movements = append(movements, currentMovement)
	}

	err = rows.Err()

	return movements, err
}

type movementStatQuery struct {
	WorkoutId uint64
	CreatedAt time.Time
	Weight    uint8
	Reps      uint8
}

func (p *PostgresRepo) GetMovementStats(movementId uint64) (types.MovementStats, error) {
	var stats types.MovementStats

	query := `
        SELECT
          workout.id AS workout_id,
          movement_set.created_at,
          movement_set.weight,
          movement_set.reps
        FROM
          movement_set
          JOIN workout_movement ON movement_set.workout_movement_id = workout_movement.id
          JOIN movement ON workout_movement.movement_id = movement.id
          JOIN workout ON workout.id = workout_movement.workout_id
        WHERE
          movement.id = $1
        ORDER BY
          created_at ASC`

	rows, err := p.db.Queryx(query, movementId)
	if err != nil {
		return stats, err
	}

	var (
		estimatedMaxes = make([]types.EstimatedMax, 0)
		bestSets       = make(map[uint64]float32)
	)

	for rows.Next() {
		var set movementStatQuery

		err = rows.Scan(&set.WorkoutId, &set.CreatedAt, &set.Weight, &set.Reps)
		currentBest, ok := bestSets[set.WorkoutId]

		estMax := maxestimate.GetMaxEstimate(float32(set.Reps), float32(set.Weight))

		if !ok {
			bestSets[set.WorkoutId] = estMax
			estimatedMaxes = append(estimatedMaxes, types.EstimatedMax{
				Date: set.CreatedAt,
				Max:  estMax,
			})
			continue
		}

		if estMax > currentBest {
			estimatedMaxes = append(estimatedMaxes, types.EstimatedMax{
				Date: set.CreatedAt,
				Max:  estMax,
			})
		}
	}

	err = rows.Err()
	if err != nil {
		return stats, err
	}

	sort.Slice(estimatedMaxes, func(i, j int) bool {
		return estimatedMaxes[i].Date.Before(estimatedMaxes[j].Date)
	})

	movement, err := p.GetMovement(movementId)
	if err != nil {
		return stats, err
	}

	currentMax := getCurrentMax(estimatedMaxes)
	allTimeMax := getAllTimeMax(estimatedMaxes)

	stats = types.MovementStats{
		EstimatedMaxes: estimatedMaxes,
		Movement:       movement,
		CurrentMax:     currentMax,
		AllTimeMax:     allTimeMax,
	}

	return stats, nil
}

func getAllTimeMax(maxes []types.EstimatedMax) types.EstimatedMax {
	var max types.EstimatedMax

	for _, m := range maxes {
		if m.Max > max.Max {
			max = m
		}
	}

	return max
}

func getCurrentMax(maxes []types.EstimatedMax) types.EstimatedMax {
	return maxes[len(maxes)-1]
}

func (p *PostgresRepo) DeleteMovement(movementId uint64) error {
	_, err := p.db.NamedExec(`DELETE FROM movement WHERE id = :movementId`,
		map[string]interface{}{
			"movementId": movementId,
		})

	return err
}

func (p *PostgresRepo) GetMovement(movementId uint64) (types.Movement, error) {
	var movement types.Movement

	row := p.db.QueryRowx("SELECT id, name, user_id FROM movement WHERE id = $1", movementId)
	err := row.Scan(&movement.ID, &movement.Name, &movement.User)

	return movement, err
}

func (p *PostgresRepo) GetMovementsByUser(userId string) ([]types.Movement, error) {
	query := `SELECT id, name 
                FROM movement
                WHERE movement.user_id = $1
                ORDER BY name ASC`

	movements := make([]types.Movement, 0)

	rows, err := p.db.Queryx(query, userId)
	if err != nil {
		return movements, err
	}

	for rows.Next() {
		var m types.Movement

		err = rows.Scan(&m.ID, &m.Name)

		movements = append(movements, m)
	}

	err = rows.Err()

	return movements, err
}
