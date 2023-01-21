package service

import "github.com/markusryoti/werk/internal/types"
import "github.com/markusryoti/werk/internal/storage"

type Service struct {
	repo storage.Storage
}

func NewService(repo storage.Storage) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddNewWorkout(userId, name string) error {
	return s.repo.AddNewWorkout(userId, name)
}

func (s *Service) GetWorkout(workoutId uint64) (types.Workout, error) {
	return s.repo.GetWorkout(workoutId)
}

func (s *Service) GetAllWorkouts() ([]types.Workout, error) {
	return s.repo.GetAllWorkouts()
}

func (s *Service) AddNewMovement(workoutId uint64, name string) error {
	return s.repo.AddNewMovement(workoutId, name)
}

func (s *Service) GetMovementsFromWorkout(workoutId uint64) ([]types.Movement, error) {
	return s.repo.GetMovementsFromWorkout(workoutId)
}

func (s *Service) AddMovementSet(movementId uint64, set types.Set) error {
	return s.repo.AddMovementSet(movementId, set)
}

func (s *Service) DeleteMovement(movementId uint64) error {
	return s.repo.DeleteMovement(movementId)
}

func (s *Service) DeleteMovementSet(movementSetId uint64) error {
	return s.repo.DeleteMovementSet(movementSetId)
}

func (s *Service) DeleteWorkout(workoutId uint64) error {
	return s.repo.DeleteWorkout(workoutId)
}
