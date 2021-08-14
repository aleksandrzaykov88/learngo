package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

//Photo describes entity which characterizes the some photo.
type Photo struct {
	ID      int    `json:"id"`
	Path    string `json:"path"`
	Name    string `json:"name"`
	Preview *Photo `json:"preview"`
}

//SetName sets name of photo-entity by its path.s
func (p *Photo) SetName() {
	s := strings.Split(p.Path, "/")
	p.Name = s[len(s)-1]
}

//Storage is an interface with methods for RESP API.
type Storage interface {
	Insert(p *Photo)
	Get() []Photo
	Delete(id int) string
}

//Gallery reserves photos and implements Storage interface.
type Gallery struct {
	counter int
	data    map[int]Photo
	sync.Mutex
}

//NewGallery constructs the MemoryStrorage object.
func NewGallery() *Gallery {
	return &Gallery{
		data:    make(map[int]Photo),
		counter: 1,
	}
}

//checkErr checks some error.
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//checkExistance checks if the item with this path exists in the database.
func checkPathExistance(db *sql.DB, path string) bool {
	rows, err := db.Query("SELECT id FROM gallery WHERE path=\"" + path + "\"")
	checkErr(err)
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}

//checkExistance checks if the item exists in the database.
func checkExistance(db *sql.DB, id string) bool {
	rows, err := db.Query("SELECT * FROM gallery WHERE id=" + id)
	checkErr(err)
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}

//Insert allows to add a new photo to the gallery.
func (g *Gallery) Insert(p *Photo) {
	db, err := sql.Open("sqlite3", Configuration.DatabasePath())
	defer db.Close()
	checkErr(err)

	if checkPathExistance(db, p.Path) {
		fmt.Println("There is already such element in database")
		return
	}

	stmt, err := db.Prepare("INSERT INTO gallery(path, name, preview) values(?,?,?)")
	checkErr(err)
	res, err := stmt.Exec(p.Path, p.Name, 0)
	checkErr(err)
	res, err = stmt.Exec(p.Preview.Path, p.Preview.Name, 1)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("Last added id is", id)
}

//Get list of photos.
func (g *Gallery) Get() []Photo {
	photos := make([]Photo, 0)
	db, err := sql.Open("sqlite3", Configuration.DatabasePath())
	defer db.Close()
	checkErr(err)

	rows, err := db.Query("SELECT id, path, name FROM gallery WHERE preview = 0")
	checkErr(err)
	var id int
	var name string
	var path string

	for rows.Next() {
		err = rows.Scan(&id, &path, &name)
		checkErr(err)
		photos = append(photos, Photo{ID: id, Name: name, Path: path})
	}

	rows.Close()
	return photos
}

//Delete photo from DB.
func (g *Gallery) Delete(id int) string {
	strId := strconv.Itoa(id)
	db, err := sql.Open("sqlite3", Configuration.DatabasePath())
	defer db.Close()
	checkErr(err)

	if !checkExistance(db, strId) {
		fmt.Println("There is no such element in database")
		return ""
	}

	rows, err := db.Query("SELECT path FROM gallery WHERE id =" + strId)
	checkErr(err)
	var path string

	//Saves path value to delete files in gallery catalog after work with db.
	for rows.Next() {
		err = rows.Scan(&path)
		checkErr(err)
	}

	rows.Close()

	stmt, err := db.Prepare("DELETE FROM gallery WHERE id=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	return path
}

//DeletePreview deletes preview from DB.
func DeletePreview(path string) {
	db, err := sql.Open("sqlite3", Configuration.DatabasePath())
	defer db.Close()
	checkErr(err)

	if !checkPathExistance(db, path) {
		fmt.Println("There is no such element in database")
	}

	stmt, err := db.Prepare("DELETE FROM gallery WHERE path=?")
	checkErr(err)

	res, err := stmt.Exec(path)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
