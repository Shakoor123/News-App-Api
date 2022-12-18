package middleware

import ("fmt"
	"net/http"
	"github.com/shakoor123/config"
	"log"

)



func GetAllUsers(w http.ResponseWriter, r *http.Request)  {
	
	type user struct {
			name string
		    age int
	}
	    rows,err:=config.DB.Query(`SELECT name,age FROM test1`)
    if err!=nil{
        log.Fatal("error ",err)
    }
    defer rows.Close()
    var users []user

    for rows.Next() {
		var u user

		err := rows.Scan( &u.name, &u.age)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", users)
	fmt.Fprint(w, "Hello World!")

	return
}
