package middleware

import(
	"net/http"
	"log"
	"github.com/shakoor123/config"
	"github.com/shakoor123/models"
	"encoding/json"
	"time"
)

func CreateNews(w http.ResponseWriter, r *http.Request){
	news:=models.News{}
	json.NewDecoder(r.Body).Decode(&news)
	t := time.Now()
	currentDate:=t.String()
	_,err:=config.DB.Query(`INSERT INTO news (uid,title,place,description,date) VALUES (?,?,?,?,?)`,news.Uid,news.Title,news.Place,news.Description,currentDate)
	if err!=nil{
		json.NewEncoder(w).Encode("Error in your request inside error")
		log.Fatal("insert news err",err)
	}else{
		json.NewEncoder(w).Encode(news)
	}
}





// t := time.Now()
// fmt.Println(t.String())
// fmt.Println(t.Format("2006-01-02 15:04:05"))