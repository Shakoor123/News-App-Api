package middleware

import (
	// "fmt"
	"net/http"
	"github.com/shakoor123/config"
	"github.com/shakoor123/models"

	"log"
	"encoding/json"

)

// type user struct {
// 	Name string `json:"name,omitempty"`
// 	Age int	`json:"age,omitempty"`
// }

func GetAllUsers(w http.ResponseWriter, r *http.Request)  {
	u:=models.User{}
	users:=[]models.User{}
	
	rows,err:=config.DB.Query(`SELECT * FROM users`)
    if err!=nil{
        log.Fatal("error ",err)
    }else{
    // defer rows.Close()

    for rows.Next() {

		 rows.Scan( &u.Uid, &u.Username,&u.Password,&u.Email)

		// if err != nil {
		// 	log.Fatal(err)
		// }
		users = append(users, u)
	}
	 json.NewEncoder(w).Encode(users)

	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%#v", users)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// response,_:=json.Marshal(users)
	// w.Write(response)
}
}
// func InsertUser(w http.ResponseWriter, r *http.Request)  {
// 	u:=models.User{}
// 	dec := json.NewDecoder(r.Body)
// 	dec.Decode(&u);
// 		_, err := config.DB.Exec(`INSERT INTO test1 (name, age) VALUES (?, ?)`, u.Name, u.Age)
// 	if err != nil {
// 		log.Fatal("not inserted",err)
// 	}else{
// 		json.NewEncoder(w).Encode("user inserted successfully")
// 	}



// }
