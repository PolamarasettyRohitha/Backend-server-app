package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	regularUserCsv = "../../utils/regularUser.csv"
	adminUserCsv   = "../../utils/adminUser.csv"
	dirPath        = "../../utils"
	username       = "username"
	password       = "password"
	admin          = "admin"
)

func home(w http.ResponseWriter, r *http.Request) {

	username := r.Header.Get(username)
	password := r.Header.Get(password)
	admin, err := strconv.ParseBool(r.Header.Get(admin))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if _, valid := isValidUser(username, password); !valid {
		http.Error(w, ErrInvalidUser.Error(), http.StatusBadRequest)
		return
	}

	if admin {
		display(w, adminUserCsv)
	}

	display(w, regularUserCsv)
}

func login(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get(username)
	password := r.URL.Query().Get(password)

	if len(username) <= 0 || len(password) <= 0 {
		http.Error(w, ErrMissingArguments.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(username, password)

	u, valid := isValidUser(username, password)
	if !valid {
		http.Error(w, ErrInvalidCredentials.Error(), http.StatusBadRequest)
		return
	}

	token, err := generatetoken(u)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Login success token : %v", token)))
}

func addBook(w http.ResponseWriter, r *http.Request) {

	username := r.Header.Get(username)
	password := r.Header.Get(password)

	if _, valid := isValidUser(username, password); !valid {
		http.Error(w, ErrInvalidUser.Error(), http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var data []string

	bookName := r.PostForm.Get("Book Name")
	author := r.PostForm.Get("Author")
	year := r.PostForm.Get("Publication Year")

	if len(bookName) <= 0 || len(author) <= 0 || len(year) <= 0 {
		http.Error(w, ErrMissingArguments.Error(), http.StatusBadRequest)
		return
	}

	val, err := strconv.Atoi(year)
	if err != nil || val > time.Now().Year() || val <= 0 {
		http.Error(w, ErrInalidArguments.Error(), http.StatusBadRequest)
		return
	}

	data = append(data, bookName, author, year)
	if err := writeCsv(regularUserCsv, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Added book successfully"))
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get(username)
	password := r.Header.Get(password)

	if _, valid := isValidUser(username, password); !valid {
		http.Error(w, ErrInvalidUser.Error(), http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	book := r.PostForm.Get("Book Name")
	if len(book) <= 0 {
		http.Error(w, ErrMissingArguments.Error(), http.StatusBadRequest)
		return
	}

	records, err := readCsv(regularUserCsv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tmpf, err := os.CreateTemp(dirPath, "00")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpf.Close()

	tmpfp := tmpf.Name()

	var found bool
	for _, eachrecord := range records {
		if strings.EqualFold(eachrecord[0], book) {
			found = true
			continue
		}

		if err := writeCsv(tmpfp, eachrecord); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	if err := renameCsv(tmpfp, regularUserCsv); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !found {
		http.Error(w, ErrNotFound.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Deleted book successfully"))
}
