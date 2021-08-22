package entities

type City struct {
	ID      int64  `json:"id"`
	StateID int64  `json:"state_id"`
	Name    string `json:"name"`
}

type State struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
