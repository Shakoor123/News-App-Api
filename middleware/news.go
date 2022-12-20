package middleware

import(
	"net/http"
	"log"
	"github.com/shakoor123/config"
	"github.com/shakoor123/models"
	"github.com/gorilla/mux"
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

func SelectAllNews(w http.ResponseWriter, r *http.Request){
	n:= models.News{}
	news:=[] models.News{}
	rows,err:=config.DB.Query("SELECT * FROM news")
	if err!=nil{
		log.Fatal("error in select news",err)
	}else{
		for rows.Next() {
			rows.Scan( &n.Nid,&n.Uid,&n.Title,&n.Place,&n.Description,&n.Date)
			news = append(news, n)
		}
		 json.NewEncoder(w).Encode(news)
	}
}

func DeleteNews(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	nid:=params["nid"]

	_,err:=config.DB.Query(`DELETE FROM news where nid=?`,nid)
	if err!=nil{
	json.NewEncoder(w).Encode("{Message:not deleted user}")
	log.Fatal("error in delete user",err)


	}else{
	json.NewEncoder(w).Encode("{Message:news deleted successfully}")
	}
}

func SelectOneNews(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	nid:=params["nid"]
	news:=models.News{}
	result,err:=config.DB.Query(`SELECT * FROM news where nid=?`,nid)
	if err!=nil{
	log.Fatal("error in select one news",err)
	}else{
		defer result.Close()
		for result.Next() {
			result.Scan( &news.Nid, &news.Uid,&news.Title,&news.Place,&news.Description,&news.Date)
		}
	}
	json.NewEncoder(w).Encode(news)
}





// t := time.Now()
// fmt.Println(t.String())
// fmt.Println(t.Format("2006-01-02 15:04:05"))