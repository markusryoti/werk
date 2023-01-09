package postgres

import (
	"time"

	"github.com/markusryoti/werk/internal/types"
)

type Workout struct {
	ID        uint64
	Date      time.Time
	Name      string
	Movements []Movement `db:"-"`
	Uid       string
}

type Movement struct {
	ID        uint64
	Name      string
	WorkoutId uint64        `db:"workout_id"`
	Sets      []MovementSet `db:"-"`
	Uid       string
}

type MovementSet struct {
	ID         uint64
	Reps       uint8
	Weight     uint8
	MovementId uint64 `db:"movement_id"`
	Uid        string
}

func (w *Workout) toDomain() types.Workout {
	sets := make([]types.Set, 0)
	movements := make([]types.Movement, 0)

	for _, mov := range w.Movements {
		for _, set := range mov.Sets {
			sets = append(sets, types.Set{
				Reps:   set.Reps,
				Weight: set.Weight,
			})
		}

		movements = append(movements, types.Movement{
			ID:   mov.ID,
			Name: mov.Name,
			Sets: nil,
		})
	}

	return types.Workout{
		ID:        w.ID,
		Date:      w.Date,
		Name:      w.Name,
		Movements: movements,
	}
}

func (m *Movement) toDomain() types.Movement {
	sets := make([]types.Set, 0)

	for _, set := range m.Sets {
		sets = append(sets, types.Set{
			Reps:   set.Reps,
			Weight: set.Weight,
		})
	}

	return types.Movement{
		ID:   m.ID,
		Name: m.Name,
		Sets: sets,
	}
}

func (s *MovementSet) toDomain() types.Set {
	return types.Set{
		Reps:   s.Reps,
		Weight: s.Weight,
	}
}
