package postgres

import "github.com/markusryoti/werk/internal/types"

func (p *PostgresRepo) AddNewWorkout(name string) error {
	_, err := p.db.NamedExec(`INSERT INTO workout (name) VALUES (:name)`,
		map[string]interface{}{
			"name": name,
		})

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepo) GetAllWorkouts() ([]types.Workout, error) {
	rows, err := p.db.Queryx(`SELECT * FROM workout`)

	workouts := make([]Workout, 0)

	for rows.Next() {
		w := Workout{}
		err = rows.StructScan(&w)
		workouts = append(workouts, w)
	}

	err = rows.Err()

	return workoutsToDomain(workouts), err
}

func workoutsToDomain(workouts []Workout) []types.Workout {
	domainWorkouts := make([]types.Workout, 0)

	for _, w := range workouts {
		domainWorkouts = append(domainWorkouts, w.toDomain())
	}

	return domainWorkouts
}
