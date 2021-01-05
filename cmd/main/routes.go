package main

import (
	"encoding/json"
	players "firstProject/basket/players/web"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(player *players.PlayerHandler) *chi.Mux {
	mux := chi.NewMux()

	// globals middleware
	mux.Use(
		middleware.Logger,    //log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	mux.Get("/hello", helloHandler)

	mux.Route("/players", func(r chi.Router) {
		r.Post("/", player.CreatePlayerHandler)
		r.Delete("/{playerId:[0-9]+}", player.DeletePlayerHandler)
		r.Patch("/{playerId:[0-9]+}", player.UpdatePlayerHandler)
	})

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "jaime")

	res := map[string]interface{}{"message": "hello world"}

	_ = json.NewEncoder(w).Encode(res)
}
