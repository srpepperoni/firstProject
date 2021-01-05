package gateway

import (
	"firstProject/basket/players/models"
	"firstProject/internal/database"
	logs "firstProject/internal/log"
	"time"
)

const MySQLTimeFormat = "2006-01-02 15:04:05"

/*
	NOTA explicativa:
	Definicion de la interfaz (en GO son implicitas: cualquier estructura que
	aplique las signaturas de sus metodos da a entender que está implementando
	la interfaz en sí)
*/

type PlayerStorage interface {
	savePlayer(cmd *models.CreatePlayerCMD) (*models.Player, error)
	editPlayer(cmd *models.UpdatePlayerCMD) (*models.Player, error)
	deletePlayer(playerId int64) *models.Player
}

type PlayerStg struct {
	// inyectamos para poder usar despues
	*database.MySqlClient
}

func (s *PlayerStg) savePlayer(cmd *models.CreatePlayerCMD) (*models.Player, error) {
	tx, err := s.Begin()

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}

	ts := time.Now().Format(MySQLTimeFormat)

	res, err := tx.Exec(`insert into player (name, last_name, height, updated_at) values (?, ?, ? , ?)`,
		cmd.Name, cmd.LastName, cmd.Height, ts)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error(err.Error())
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

func (s *PlayerStg) editPlayer(cmd *models.UpdatePlayerCMD) (*models.Player, error) {
	tx, err := s.Begin()

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}

	ts := time.Now().Format(MySQLTimeFormat)

	res, err := tx.Exec(`update player set name = ?, last_name = ?, height = ?, updated_at = ? where id = ?`,
		cmd.Name, cmd.LastName, cmd.Height, ts, cmd.ID)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error(err.Error())
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

func (s *PlayerStg) deletePlayer(playerId int64) *models.Player {
	tx, err := s.Begin()

	if err != nil {
		logs.Log().Error(err.Error())
		return nil
	}

	var player models.Player
	err = tx.QueryRow(`select name, last_name, height from player where id = ?`, playerId).Scan(&player.Name, &player.LastName, &player.Height)

	if err != nil {
		_ = tx.Rollback()
		return nil
	}

	_, err = tx.Exec(`delete from player where id = ?`, playerId)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil
	}

	_ = tx.Commit()
	return &models.Player{
		ID:       playerId,
		Name:     player.Name,
		LastName: player.LastName,
		Height:   player.Height,
	}
}
