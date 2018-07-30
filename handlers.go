package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Get article from URL parameter
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
		log.Fatal(err)
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

// Create article from POST request.
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
	err = t.Check()
	if err != nil {
		render.Render(w, r, ErrRequiredField(err))
		return
	}

	t.Date = time.Now().Format("2006-01-02")
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
				log.Fatal(err)
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
		log.Fatal(err)
		render.Render(w, r, ErrServerSide(err))
		return
	}
	newArticleId, _ = res.LastInsertId()

	for _, tid := range tagIds {
		_, err = tx.Exec("INSERT INTO tags_for_articles (article_id, tag_id) VALUES (?, ?)", newArticleId, tid)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			render.Render(w, r, ErrServerSide(err))
			return
		}
	}

	// Presumably if we've managed to get here then the queries were fine.
	t.Id = int(newArticleId)
	tx.Commit()
	render.JSON(w, r, t)

}

// Get tagged articles from url parameters.
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
		"WHERE b.title = ? AND c.date = ? "+
		"LIMIT 10", title, dateFm)

	if err != nil {
		log.Fatal(err)
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
	// This is the string substitution I mention in the wiki.
	// It doesn't feel right, but I've spoken to other gophers who
	// say that as long as the injection is not possible it's fine.
	sql := fmt.Sprintf("SELECT DISTINCT tags.title FROM tags_for_articles "+
		"LEFT JOIN tags ON tags.id = tags_for_articles.tag_id "+
		"LEFT JOIN articles ON tags_for_articles.tag_id = articles.id "+
		"WHERE date = \"%s\" AND articles.id IN (%s)", dateFm, inq)
	fmt.Println(sql)
	rows, err = db.Query(sql)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&tagTitle)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tagTitle)
		tags = append(tags, tagTitle)
	}
	rows.Close()

	responseObj := new(Tags)
	responseObj.Count = len(tags)
	responseObj.Title = title
	responseObj.Articles = artIds
	responseObj.RelatedTags = tags

	render.JSON(w, r, responseObj)
}
