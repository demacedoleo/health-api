package location

import "context"

type Repository interface {
	GetCities(ctx context.Context, stateID int64) ([]City, error)
	GetStates(ctx context.Context) ([]State, error)
}

type City struct {
	ID      int64
	StateID int64
	Name    string
}

type State struct {
	ID   int64
	Name string
}
