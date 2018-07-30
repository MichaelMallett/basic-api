package main

import (
	"database/sql"
	"encoding/json"
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

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func ErrServerSide(err error) render.Renderer {
	log.Fatal(err)
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Server Error",
		ErrorText:      err.Error(),
	}
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRequiredField(name string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: 400,
		StatusText:     "Missing field",
		ErrorText:      name + " is a required field.",
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

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
	defer stmt.Close()
	if err != nil {
		render.Render(w, r, ErrServerSide(err))
		return
	}
	err = stmt.QueryRow(id).Scan(&tagString, &articleId, &title, &body, &date)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	// I could get this in a seperate query straight to an array, but I have always
	// tried to 'get the database to do the work', so I err towards joins. In production
	// I would probably benchmark something like this and see what's better maybe?
	tags := strings.Split(tagString, ",")

	// I don't know if there's a better way to do this, but this seems a bit messy.
	// Get datetime from sql, parse it to time object, format to string.
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatal(err)
	}
	fD := t.Format("02-01-2006")

	responseObj := new(Article)
	responseObj.Id = articleId
	responseObj.Title = title
	responseObj.Body = body
	responseObj.Tags = tags
	responseObj.Date = fD

	render.JSON(w, r, responseObj)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	var t Article
	var tagId int
	var tagIds []int
	var tagTitle string
	var insertTagId int64
	var newArticleId int64

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	t.Date = time.Now().Format("2006-01-02")
	fmt.Printf("%T", t.Date)

	// Deal with new tags first.
	// I'm really not that keen on doing so many select queries, ideally would cache
	// a list of tags to check against, perhaps a static file, but this was the lesser of
	// the evils I tried.
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	for _, ti := range t.Tags {
		db.QueryRow("SELECT id, title FROM tags WHERE title = ?", ti).Scan(&tagId, &tagTitle)
		if ti == tagTitle {
			tagIds = append(tagIds, tagId)
		} else {
			res, err := tx.Exec("INSERT INTO tags (title) VALUES (?)", ti)
			if err != nil {
				tx.Rollback()
				render.Render(w, r, ErrServerSide(err))
				return
			}
			// Add to tags slice, we need this later.
			insertTagId, _ = res.LastInsertId()
			// I do a type conversion here, which I think is probably a bad thing to do.
			tagIds = append(tagIds, int(insertTagId))
		}
	}

	res, err := tx.Exec("INSERT INTO articles (title, body, date) VALUES (?, ?, ?)", t.Title, t.Body, t.Date)
	if err != nil {
		tx.Rollback()
		render.Render(w, r, ErrServerSide(err))
		return
	}
	newArticleId, _ = res.LastInsertId()

	for _, tid := range tagIds {
		_, err = tx.Exec("INSERT INTO tags_for_articles (article_id, tag_id) VALUES (?, ?)", newArticleId, tid)
		if err != nil {
			tx.Rollback()
			render.Render(w, r, ErrServerSide(err))
			return
		}
	}

	// Presumably if we've managed to get here then the queries were fine.
	t.Id = int(newArticleId)
	tx.Commit()
	render.JSON(w, r, t)

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

	fmt.Println(title)
	fmt.Println(dateFm)
	rows, err := db.Query("SELECT article_id FROM tags_for_articles a "+
		"LEFT JOIN tags b ON a.tag_id = b.id "+
		"LEFT JOIN articles c ON a.article_id = c.id "+
		"WHERE b.title = ? AND c.date = ?", title, dateFm)

	if err != nil {
		render.Render(w, r, ErrServerSide(err))
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
		render.Render(w, r, ErrNotFound)
		return
	}
	inq := strings.Join(artIds, ",")
	sql := fmt.Sprintf("SELECT DISTINCT title FROM tags_for_articles LEFT JOIN tags ON tags.id = tags_for_articles.tag_id WHERE article_id IN (%s)", inq)
	rows, err = db.Query(sql)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&tagTitle)
		// fmt.Println(tagTitle)
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

	render.JSON(w, r, responseObj)
}
