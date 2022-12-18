package router

import(
	"github.com/shakoor123/middleware"
	"github.com/gorilla/mux"

)


func Router() *mux.Router{
	router:=mux.NewRouter()

	router.HandleFunc("/api/users",middleware.GetAllUsers).Methods("GET","OPTIONS")
	
	return router
}