package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Article struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
}

type Tags struct {
	Title       string   `json:"tag"`
	Count       int      `json:"count"`
	Articles    []string `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}

// Global database connection.
var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3308)/fairfax?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Error")
	}

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

func getArticle(w http.ResponseWriter, r *http.Request) {

	var (
		tagString string
		articleId int
		title     string
		body      string
		date      string
	)

	id := chi.URLParam(r, "articleId")
	stmt, err := db.Prepare("SELECT group_concat(a.title), b.* FROM tags a " +
		"JOIN tags_for_articles ON a.id = tags_for_articles.tag_id " +
		"JOIN articles b ON tags_for_articles.article_id = b.id " +
		"WHERE b.id = ? " +
		"GROUP BY b.id")
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.QueryRow(id).Scan(&tagString, &articleId, &title, &body, &date)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// I could get this in a seperate query straight to an array, but I have always
	// tried to 'get the database to do the work', so I err towards joins. In production
	// I would probably benchmark something like this.
	tags := strings.Split(tagString, ",")

	// I don't know if there's a better way to do this, but this seems a bit messy.
	// Get datetime from sql, parse it to time object, format to string.
	t, err := time.Parse(time.RFC3339, date)
	fD := t.Format("02-03-2006")

	responseObj := new(Article)
	responseObj.Id = articleId
	responseObj.Title = title
	responseObj.Body = body
	responseObj.Tags = tags
	responseObj.Date = fD

	render.JSON(w, r, responseObj)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Article"))

}

func getTaggedArticles(w http.ResponseWriter, r *http.Request) {

	var artIds []string
	var artId int
	var tags []string
	var tagTitle string

	title := chi.URLParam(r, "tagName")
	date := chi.URLParam(r, "date")
	dateObj, err := time.Parse("20060102", date)
	if err != nil {
		log.Fatal(err)
	}
	dateFm := dateObj.Format("2006-01-02")

	rows, err := db.Query("SELECT article_id FROM tags_for_articles a "+
		"LEFT JOIN tags b ON a.tag_id = b.id "+
		"LEFT JOIN articles c ON a.article_id = c.id "+
		"WHERE b.title = ? AND c.date = ?", title, dateFm)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&artId)
		if err != nil {
			log.Fatal(err)
		}
		artIds = append(artIds, strconv.Itoa(artId))
	}
	rows.Close()

	if len(artIds) == 0 {
		log.Fatal(err)
	}

	rows, err = db.Query("SELECT DISTINCT title FROM tags_for_articles "+
		"LEFT JOIN tags ON tags.id = tags_for_articles.tag_id "+
		"WHERE article_id IN (?)", strings.Join(artIds, ","))
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&tagTitle)
		if err != nil {
			log.Fatal(err)
		}
		tags = append(tags, tagTitle)
	}
	rows.Close()

	responseObj := new(Tags)
	responseObj.Count = len(artIds)
	responseObj.Title = title
	responseObj.Articles = artIds
	responseObj.RelatedTags = tags

	w.Write([]byte("Get Tagged Articles"))
	render.JSON(w, r, responseObj)
}
