package main

import (
	"fmt"
	"html/template"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id                            uint16
	Title, Announcement, FullText string
}

var posts = []Article{}
var showPost = Article{}

//index handles index page of website.
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT *  FROM `articles`")
	if err != nil {
		panic(err)
	}
	posts = []Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Announcement, &post.FullText)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	t.ExecuteTemplate(w, "index", posts)
}

//create builds page with input form.
func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

//save_article saves text into database.
func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	announcement := r.FormValue("announcement")
	full_text := r.FormValue("full_text")

	if title == "" || announcement == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {

		db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `announcement`, `full_text`) VALUES ('%s', '%s', '%s')", title, announcement, full_text))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

//show_post shows article on web page.
func show_post(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	vars := mux.Vars(r)
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT *  FROM `articles` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}
	showPost = Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Announcement, &post.FullText)
		if err != nil {
			panic(err)
		}
		showPost = post
	}

	t.ExecuteTemplate(w, "show", showPost)
}

func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET", "POST")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
