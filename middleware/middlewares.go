package middleware

import ("fmt"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello World!")
	return
}
