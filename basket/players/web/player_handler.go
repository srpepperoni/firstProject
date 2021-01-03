package web

import (
	"encoding/json"
	"firstProject/basket/players/gateway"
	"firstProject/basket/players/models"
	"firstProject/internal/database"
	"log"
	"net/http"
)

type PlayerHandler struct {
	gateway.PlayerGateway
}

func NewPlayerHandler(client *database.MySqlClient) *PlayerHandler {
	return &PlayerHandler{gateway.NewPlayerGateway(client)}
}

func (h *PlayerHandler) CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd, err := parseCreateRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res, err := h.SavePlayer(cmd)

	if err != nil {
		log.Fatalln(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot save player"})
		return
	}

	json.NewEncoder(w).Encode(&res)

}

func parseCreateRequest(r *http.Request) (*models.CreatePlayerCMD, error) {
	body := r.Body
	defer body.Close()
	var cmd models.CreatePlayerCMD
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return &cmd, nil
}
