package router

import(
	"github.com/shakoor123/middleware"
	"github.com/gorilla/mux"

)


func Router() *mux.Router{
	router:=mux.NewRouter()

	router.HandleFunc("/api/auth/register",middleware.RegisterUSer).Methods("POST","OPTIONS")
	router.HandleFunc("/api/auth/login",middleware.LoginUser).Methods("POST","OPTIONS")
	router.HandleFunc("/api/users",middleware.GetAllUsers).Methods("GET","OPTIONS")
	router.HandleFunc("/api/news",middleware.CreateNews).Methods("POST","OPTIONS")

	
	return router
}