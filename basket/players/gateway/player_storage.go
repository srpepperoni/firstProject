package gateway

import (
	"firstProject/basket/players/models"
	"firstProject/internal/database"
	"log"
	"time"
)

/*
	NOTA explicativa:
	Definicion de la interfaz (en GO son implicitas: cualquier estructura que
	aplique las signaturas de sus metodos da a entender que está implementando
	la interfaz en sí)
*/

type PlayerStorage interface {
	savePlayer(cmd *models.CreatePlayerCMD) (*models.Player, error)
	editPlayer(cmd *models.UpdatePlayerCMD) (*models.Player, error) // TODO
	deletePlayer(playerId int64) (*models.Player, error) // TODO
}

type PlayerStg struct {
	// inyectamos para poder usar despues
	*database.MySqlClient
}

func (s *PlayerStg) savePlayer(cmd *models.CreatePlayerCMD) (*models.Player, error) {
	tx, err := s.Begin()

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	ts := time.Now().Unix()

	res, err := tx.Exec(`insert into player (name, last_name, height, updated_at) values (?, ?, ? , ?)`,
		cmd.Name, cmd.LastName, cmd.Height, ts)

	if err != nil {
		log.Fatalln(err.Error())
		_ = tx.Rollback()
		return nil, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		log.Fatalln(err.Error())
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()
	return &models.Player{
		ID:       lastID,
		Name:     cmd.Name,
		LastName: cmd.LastName,
		Height:   cmd.Height,
	}, nil
}
