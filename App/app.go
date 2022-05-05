package App

import (
	"TimeAPI/Routes/Time"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer() {
	router := mux.NewRouter()

	timeRouter := router.PathPrefix("/api").Subrouter()
	Time.RegisterTimeRoutes(timeRouter)

	_ = http.ListenAndServe("localhost:3000", router)
}
