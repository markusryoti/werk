package service

import "github.com/markusryoti/werk/internal/types"

type Service interface {
	AddNewWorkout(userId, name string) error
	GetWorkout(workoutId uint64) (types.Workout, error)
	GetAllWorkouts() ([]types.Workout, error)
	AddNewMovement(workoutId uint64, name string) error
	GetMovementsFromWorkout(workoutId uint64) ([]types.Movement, error)
	AddMovementSet(movementId uint64, set types.Set) error
}
