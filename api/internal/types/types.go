package types

import "time"

type Workout struct {
	ID        uint64     `json:"id"`
	Date      time.Time  `json:"date"`
	Name      string     `json:"name"`
	Movements []Movement `json:"movements"`
	User      string     `json:"user"`
}

type Movement struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Sets []Set  `json:"sets"`
	User string `json:"-"`
}

type Set struct {
	ID     uint64 `json:"id"`
	Reps   uint8  `json:"reps"`
	Weight int    `json:"weight"`
	User   string `json:"-"`
}

type MovementStats struct {
	Movement       Movement       `json:"movement"`
	EstimatedMaxes []EstimatedMax `json:"estimatedMaxes"`
}

type EstimatedMax struct {
	Date time.Time `json:"date"`
	Max  float32   `json:"max"`
}
