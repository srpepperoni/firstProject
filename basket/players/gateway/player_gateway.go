package gateway

import (
	"firstProject/basket/players/models"
	"firstProject/internal/database"
)

// La eleccion de usar interfaces es por temas de testing a la hora de hacer mocks (entre otros)

type PlayerGateway interface {
	SavePlayer(cmd *models.CreatePlayerCMD) (*models.Player, error)
	DeletePlayer(playerId int64) *models.Player
	UpdatePlayer(cmd *models.UpdatePlayerCMD) (*models.Player, error)
}

type PlayerGtw struct {
	PlayerStorage
}

func NewPlayerGateway(client *database.MySqlClient) PlayerGateway {
	return &PlayerGtw{&PlayerStg{client}}
}

// SavePlayer stores a player in the DB
func (g *PlayerGtw) SavePlayer(cmd *models.CreatePlayerCMD) (*models.Player, error) {
	return g.savePlayer(cmd)
}

// DeletePlayer delete the player in the DB, with an ID
func (g *PlayerGtw) DeletePlayer(playerId int64) *models.Player {
	return g.deletePlayer(playerId)
}

// UpdatePlayer updates the player in the DB
func (g *PlayerGtw) UpdatePlayer(cmd *models.UpdatePlayerCMD) (*models.Player, error) {
	return g.editPlayer(cmd)
}


