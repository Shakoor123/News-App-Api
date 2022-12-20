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
	router.HandleFunc("/api/users/{id}",middleware.SelectOneUser).Methods("GET","OPTIONS")
	router.HandleFunc("/api/users/{id}",middleware.DeleteUser).Methods("DELETE","OPTIONS")

	router.HandleFunc("/api/news",middleware.CreateNews).Methods("POST","OPTIONS")
	router.HandleFunc("/api/news",middleware.SelectAllNews).Methods("GET","OPTIONS")
	router.HandleFunc("/api/news/{nid}",middleware.DeleteNews).Methods("DELETE","OPTIONS")
	router.HandleFunc("/api/news/{nid}",middleware.SelectOneNews).Methods("GET","OPTIONS")
	
	return router
}