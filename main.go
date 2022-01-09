package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home end point hit")
	w.Write([]byte("<h1>Home end point hit , welcome!</h1>"))
}

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var articles []Article

func (a *Article) IsEmpty() bool {
	return a.Title == ""
}
func main() {

	handleRequest()
}
func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/all", getAllArticles)
	r.HandleFunc("/create",createArticle)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles = append(articles, Article{Id: "0", Title: "Hello", Desc: "Article Description", Content: "Article Content"})
	articles = append(articles, Article{Id: "1", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"})

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func getOneArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	articleId := params["id"]
	for _, article := range articles {
		if articleId == article.Id {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	json.NewEncoder(w).Encode("No article with that id found")
	return
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data to the server")
	}
	// if empty body {}
	var article Article
	json.NewDecoder(r.Body).Decode(&article)
	if article.IsEmpty() {
		json.NewEncoder(w).Encode("No data found")
		return
	}
	//with data , generate id for the post
	rand.Seed(time.Now().UnixNano())
	article.Id = strconv.Itoa(rand.Intn(100))
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
	return

}

func updateArticle (w http.ResponseWriter , r *http.Request){
	
}