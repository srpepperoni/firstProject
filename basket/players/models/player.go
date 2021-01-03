package models

type Player struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Height   float64 `json:"height"`
}

type CreatePlayerCMD struct {
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Height   float64 `json:"height"`
}

type UpdatePlayerCMD struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Height   float64 `json:"height"`
}
