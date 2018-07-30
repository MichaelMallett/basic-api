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

// Global database connection.
var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql",
		"root:password@tcp(db:3306)/fairfax?parseTime=true")
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
		r.Post("/", createArticle)
	})

	r.Route("/tags", func(r chi.Router) {
		r.Get("/{tagName}/{date}", getTaggedArticles)
	})

	http.ListenAndServe(":3000", r)

}
