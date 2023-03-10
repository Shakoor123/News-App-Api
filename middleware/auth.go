package middleware

import (
	// "fmt"
	"net/http"
	"github.com/shakoor123/config"
	"github.com/shakoor123/models"

	"log"
	"encoding/json"

)

func LoginUser(w http.ResponseWriter,r *http.Request){
	user:=models.User{}
	u:=models.User{}
	dec := json.NewDecoder(r.Body)
	dec.Decode(&u);
	result, err := config.DB.Query(`select * from users where email=?`,u.Email)
	if err!=nil{
		log.Fatal(err)
	}
	defer result.Close()
	for result.Next() {
	result.Scan( &user.Uid, &user.Username,&user.Password,&user.Email)
	}

	// if user.Password == u.Password{
	// json.NewEncoder(w).Encode(user)

	// }else{
	// json.NewEncoder(w).Encode("Password is incurrect")
	// }
	flag:=CheckPasswordHash(u.Password,user.Password)
	if flag==true{
		session,_:=Store.Get(r,"session")
		session.Values["userId"]=user.Uid;
		session.Save(r,w)
		json.NewEncoder(w).Encode(user)
	}else{
		json.NewEncoder(w).Encode("Password is incurrect")
	}

}

func RegisterUSer(w http.ResponseWriter,r *http.Request){
	u:=models.User{}
	dec := json.NewDecoder(r.Body)
	dec.Decode(&u);
	hash,_:=HashPassword(u.Password)
	_, err := config.DB.Exec(`INSERT INTO users (username,password,email) VALUES (?, ?,?)`, u.Username,hash,u.Email)
	if err != nil {
		log.Fatal("not inserted",err)
	}else{
		json.NewEncoder(w).Encode("user inserted successfully")
	}
}