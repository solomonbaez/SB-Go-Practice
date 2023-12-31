package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var e error
	db, e = sql.Open("mysql", cfg.FormatDSN())
	if e != nil {
		log.Fatal(e)
	}

	ping := db.Ping() // confirm database connection
	if ping != nil {
		log.Fatal(ping)
	}

	fmt.Println("Connected to MySQL!")

	var albums []Album
	albums, e = albumsByArtist(("Amiture"))
	if e != nil {
		log.Fatal(e)
	}

	aId, e := addAlbum(Album{
		Title:  "Substance",
		Artist: "New Order",
		Price:  14.49,
	})
	if e != nil {
		log.Fatal(e)
	}

	var a Album
	a, e = albumByID(aId)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("Album added: %v\n", a)

	e = deleteAlbum(aId)
	if e != nil {
		log.Fatal(e)
	}

	e = deleteAlbum(aId)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("Album deleted: %v\n", a)

	fmt.Printf("Albums found: %v\n", albums)
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	var e error

	rows, e := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if e != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, e)
	}

	defer rows.Close() // defer row backoff
	for rows.Next() {
		var a Album
		if e := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); e != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, e)
		}

		albums = append(albums, a)
	}

	if e := rows.Err(); e != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, e)
	}

	return albums, nil
}

func albumByID(id int64) (Album, error) {
	var a Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if e := row.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); e != nil {
		if e == sql.ErrNoRows {
			return a, fmt.Errorf("albumsByID %d: no such album", id)
		}
		return a, fmt.Errorf("albumsByID %d: %v", id, e)
	}

	return a, nil
}

func addAlbum(a Album) (int64, error) {
	var e error
	result, e := db.Exec(
		"INSERT INTO album (title, artist, price) VALUES (?, ?, ?)",
		a.Title,
		a.Artist,
		a.Price,
	)
	if e != nil {
		return 0, fmt.Errorf("addAlbum: %v", e)
	}

	id, e := result.LastInsertId()
	if e != nil {
		return 0, fmt.Errorf("addAlbum: %v", e)
	}

	return id, nil
}

func deleteAlbum(id int64) error {
	var e error
	result, e := db.Exec(
		"DELETE FROM album WHERE id=?",
		id,
	)
	if e != nil {
		return fmt.Errorf("deleteAlbum %d: %v", result, e)
	}

	return nil
}
