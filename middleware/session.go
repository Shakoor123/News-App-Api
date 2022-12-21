package middleware

import(
	// "fmt"
	"github.com/gorilla/sessions"
	"net/http"


)

var Store = sessions.NewCookieStore([]byte("super-secret"))


func Auth(w http.ResponseWriter, r *http.Request)bool {
	session, _ := Store.Get(r, "session")
	_, ok := session.Values["userId"]
	if !ok {
	return false
	}
	return true
}