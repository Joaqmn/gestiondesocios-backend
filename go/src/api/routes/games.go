package routes

import (
	"fs5-gestiondesocios-backend/api/controllers"

	"github.com/gorilla/mux"
)

//SetGamesRoutes sets the game routes
func SetGamesRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/games/{id}",
		controllers.GetGame).Methods("GET")
	subRouter.HandleFunc("/games",
		controllers.GetGames).Methods("GET")
	subRouter.HandleFunc("/games",
		controllers.InsertGame).Methods("POST")
	subRouter.HandleFunc("/games/{id}",
		controllers.DeleteGame).Methods("DELETE")
	subRouter.HandleFunc("/games/{id}",
		controllers.UpdateGame).Methods("PUT")
}
