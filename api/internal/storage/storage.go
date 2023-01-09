package storage

type Storage interface {
	AddNewWorkout(name string) error
}
