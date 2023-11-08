package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type RequestGet struct {
		Title   string
		Message string
	}

	var requestget RequestGet

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
	r.Body.Close()

	errors := json.Unmarshal(body, &requestget)
	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	}

	if requestget.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Both 'Title' fields are required")
		return
	}

	row := db.QueryRow("select * from text where title  = $1", requestget.Title)
	responsen := RequestGet{}
	err = row.Scan(&requestget.Title, &requestget.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}

	response, err := json.Marshal(responsen)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func put(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type RequestPut struct {
		Title   string
		Message string
	}
	var requestput RequestPut

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
	r.Body.Close()

	errors := json.Unmarshal(body, &requestput)
	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	}

	if requestput.Title == "" || requestput.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Both 'Title' and 'Message' fields are required")
		return
	}

	_, err = db.Exec("INSERT INTO text (title,message) VALUES (?, ?)", requestput.Title, requestput.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}

}

func deleted(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type RequestDeleted struct {
		Title string
	}

	var requestdeleted RequestDeleted

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
	r.Body.Close()

	errors := json.Unmarshal(body, &requestdeleted)
	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	}

	_, err = db.Exec("DELETE FROM text WHERE title = ?", requestdeleted.Title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}

}

func post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type RequestPost struct {
		Title   string
		Message string
	}

	var requestpost RequestPost

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
	r.Body.Close()

	errors := json.Unmarshal(body, &requestpost)
	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	}

	if requestpost.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty element is not allowed")
		return
	}

	_, err = db.Exec("INSERT INTO text (title, message) VALUES (?, ?)", requestpost.Title, requestpost.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func main() {
	var err error

	db, err = sql.Open("sqlite3", "store.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS text (
			id INTEGER PRIMARY KEY,
			title TEXT,
			message TEXT
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.GET("/get", get)
	router.PUT("/put", put)
	router.DELETE("/delete", deleted)
	router.POST("/post", post)
	http.ListenAndServe(":8080", router)
}
