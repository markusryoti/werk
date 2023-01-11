package storage

import "github.com/markusryoti/werk/internal/types"

type Storage interface {
	AddNewWorkout(name string) error
	AddMovementSet(movementId uint64, set types.Set) error
}
