package router

import (
	"github.com/Haidar1528/mongodb_mux_golang/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("api/movie/{id}", controller.DeleteOneMoive).Methods("DELETE")
	router.HandleFunc("api/deleteallmovies", controller.DeleteAllMoives).Methods("DELETE")

	return router
}
