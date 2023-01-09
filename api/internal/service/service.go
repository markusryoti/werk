package service

import "github.com/markusryoti/werk/internal/types"

type Repository interface {
	AddNewWorkout(name string) error
	GetAllWorkouts() ([]types.Workout, error)
	AddNewMovement(workoutId uint64, name string) error
	GetMovementsFromWorkout(workoutId uint64) ([]types.Movement, error)
}
