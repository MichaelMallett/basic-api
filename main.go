package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Article struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
}

// Global database connection.
var db *sql.DB

func main() {

	db, err := sql.Open("mysql",
		"fairfax:password@tcp(db:3306)/fairfax")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/articles", func(r chi.Router) {
		r.Get("/{articleId}", getArticle)
		r.Get("/", createArticle)
	})

	r.Route("/tags", func(r chi.Router) {
		r.Get("/{tagName}/{date}", getTaggedArticles)
	})

	http.ListenAndServe(":3000", r)

}

func getArticle(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "articleId")
	stmt, err := db.Prepare("SELECT title, body, date FROM articles WHERE id=?")
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Article"))

}

func getTaggedArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Tagged Articles"))

}
