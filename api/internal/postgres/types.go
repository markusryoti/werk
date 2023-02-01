package postgres

import (
	"time"

	"github.com/markusryoti/werk/internal/types"
)

type Workout struct {
	ID        uint64
	Name      string    `db:"name"`
	UserId    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Movement struct {
	ID        uint64
	Name      string    `db:"name"`
	UserId    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type WorkoutMovement struct {
	ID         uint64
	WorkoutId  uint64    `db:"workout_id"`
	MovementId uint64    `db:"movement_id"`
	UserId     string    `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type MovementSet struct {
	ID                uint64
	Reps              uint8     `db:"reps"`
	Weight            int       `db:"weight"`
	WorkoutMovementId uint64    `db:"workout_movement_id"`
	UserId            string    `db:"user_id"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

func (w *Workout) toDomain() types.Workout {
	return types.Workout{
		ID:   w.ID,
		Name: w.Name,
		User: w.UserId,
		Date: w.CreatedAt,
	}
}
