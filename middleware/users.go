package middleware

import (
	// "fmt"
	"net/http"
	"github.com/shakoor123/config"
	"github.com/shakoor123/models"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"

)

func GetAllUsers(w http.ResponseWriter, r *http.Request)  {
	flag:=Auth(w,r)
	if !flag{
		json.NewEncoder(w).Encode("not authenticated")
		return
	}

	u:=models.User{}
	users:=[]models.User{}
	
	rows,err:=config.DB.Query(`SELECT * FROM users`)
    if err!=nil{
        log.Fatal("error ",err)
    }else{
    for rows.Next() {
		rows.Scan( &u.Uid, &u.Username,&u.Password,&u.Email)
		users = append(users, u)
	}
	 json.NewEncoder(w).Encode(users)
	// fmt.Printf("%#v", users)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// response,_:=json.Marshal(users)
	// w.Write(response)
}
}


func DeleteUser(w http.ResponseWriter,r *http.Request){
	flag:=Auth(w,r)
	if !flag{
		json.NewEncoder(w).Encode("not authenticated")
		return
	}
	params:=mux.Vars(r)
	id:=params["id"]

	_,err:=config.DB.Query(`DELETE FROM users where uid=?`,id)
	if err!=nil{
	json.NewEncoder(w).Encode("{Message:not deleted user}")
	log.Fatal("error in delete user",err)


	}else{
	json.NewEncoder(w).Encode("{Message:user deleted successfully}")
	}
}

func SelectOneUser(w http.ResponseWriter,r *http.Request){
	flag:=Auth(w,r)
	if !flag{
		json.NewEncoder(w).Encode("not authenticated")
		return
	}
	params:=mux.Vars(r)
	id:=params["id"]
	user:=models.User{}
	result,err:=config.DB.Query(`SELECT * FROM users where uid=?`,id)
	if err!=nil{
	json.NewEncoder(w).Encode("{Message:not deleted user}")
	log.Fatal("error in delete user",err)
	}else{
		defer result.Close()
		for result.Next() {
			result.Scan( &user.Uid, &user.Username,&user.Password,&user.Email)
		}
	}
	json.NewEncoder(w).Encode(user)
}
