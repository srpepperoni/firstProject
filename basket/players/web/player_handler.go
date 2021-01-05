package web

import (
	"encoding/json"
	"firstProject/basket/players/gateway"
	"firstProject/basket/players/models"
	"firstProject/internal/database"
	logs "firstProject/internal/log"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
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
		logs.Log().Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot save player"})
		return
	}

	json.NewEncoder(w).Encode(&res)

}

func (h *PlayerHandler) DeletePlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	playerId, err := strconv.ParseInt(chi.URLParam(r, "playerId"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res := h.DeletePlayer(playerId)

	if res == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot delete player"})
		return
	}

	json.NewEncoder(w).Encode(&res)
}

func (h *PlayerHandler) UpdatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	playerId, err := strconv.ParseInt(chi.URLParam(r, "playerId"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters from request"})
		return
	}

	cmd, err := parseUpdateRequest(r, playerId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res, err := h.UpdatePlayer(cmd)

	json.NewEncoder(w).Encode(&res)
}

func parseCreateRequest(r *http.Request) (*models.CreatePlayerCMD, error) {
	body := r.Body
	defer body.Close()
	var cmd models.CreatePlayerCMD
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}

	return &cmd, nil
}

func parseUpdateRequest(r *http.Request, id int64) (*models.UpdatePlayerCMD, error) {
	body := r.Body
	defer body.Close()
	var cmd models.UpdatePlayerCMD
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}
	cmd.ID = id
	return &cmd, nil
}
