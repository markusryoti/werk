package storage

import "github.com/markusryoti/werk/internal/types"

type Storage interface {
	AddNewWorkout(userId, name string) error
	GetWorkout(workoutId uint64) (types.Workout, error)
	GetMovement(movementId uint64) (types.Movement, error)
	GetMovementSet(movementSetId uint64) (types.Set, error)
	GetAllWorkouts() ([]types.Workout, error)
	AddNewMovement(workoutId uint64, movement types.Movement) error
	GetMovementsFromWorkout(workoutId uint64) ([]types.Movement, error)
	AddMovementSet(movementId uint64, set types.Set) error
	DeleteWorkout(workoutId uint64) error
	DeleteMovement(movementId uint64) error
	DeleteMovementSet(movementSetId uint64) error
}
