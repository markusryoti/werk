package types

import "time"

type Workout struct {
	ID        uint64     `json:"id"`
	Date      time.Time  `json:"date"`
	Name      string     `json:"name"`
	Movements []Movement `json:"movements"`
	User      string
}

type Movement struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Sets []Set  `json:"sets"`
}

type Set struct {
	Reps   uint8 `json:"reps"`
	Weight uint8 `json:"weight"`
}

type User struct {
	Id        string
	Username  string
	FirstName string
	LastName  string
}
