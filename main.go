package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

type Article struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
}

func main() {
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
	w.Write([]byte("Get Article Method"))
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Article"))

}

func getTaggedArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Tagged Articles"))

}
