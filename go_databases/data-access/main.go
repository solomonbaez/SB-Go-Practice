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

	fmt.Printf("Albums found: %v\n", albums)
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

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
