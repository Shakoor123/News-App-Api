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