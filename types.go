package main

import (
	"github.com/go-chi/render"
	"log"
	"net/http"
)

// Content models.
type Article struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
}

// Error checking on Article.
func (a *Article) Check() error {
	if len(a.Title) == 0 {
		return ErrMissingField("title")
	}
	if len(a.Body) == 0 {
		return ErrMissingField("body")
	}
	if len(a.Tags) == 0 {
		return ErrMissingField("tags")
	}
	return nil
}

type Tags struct {
	Title       string   `json:"tag"`
	Count       int      `json:"count"`
	Articles    []string `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}

// Error Responses.
type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	AppCode        int64  `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
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

func ErrRequiredField(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Missing field",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

// Use Error interface for Missing field.
type ErrMissingField string

func (e ErrMissingField) Error() string {
	return string(e) + " is required."
}
