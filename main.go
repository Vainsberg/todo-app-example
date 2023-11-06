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

type Request struct {
	Message string
}

func get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method Not Allowed")
		return
	}
	var requestget Request

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
	if requestget.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty element is not allowed")
		return
	}

	_, err = db.Exec("SELECT * FROM text WHERE name = ?", requestget.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
}

func put(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method Not Allowed")
		return
	}
	var requestget Request
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

	if requestget.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty element is not allowed")
		return
	}

	_, err = db.Exec("INSERT INTO text (name) VALUES (?)", requestget.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
}

func deleted(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method Not Allowed")
		return
	}
	var requestget Request
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
	_, err = db.Exec("DELETE FROM text WHERE name = ?", requestget.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
}

func post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method Not Allowed")
		return
	}
	var requestget Request
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
	if requestget.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty element is not allowed")
		return
	}

	_, err = db.Exec("INSERT INTO text (name) VALUES (?)", requestget.Message)
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
			name TEXT
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
