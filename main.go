package main

import(
	"fmt"
	"net/http"
	"log"
	"github.com/shakoor123/router"
	
)

func main(){
	r:=router.Router()

	fmt.Println("Listening on 4000 ")
	log.Fatal(http.ListenAndServe(":4000",r))
}

// https://github.com/boantp/go-mysql-rest-api/blob/master/controllers/web.go

// https://github.com/schadokar/go-to-do-app/blob/main/go-server/middleware/middleware.go