package main

import (
	"database/sql"
	"fmt"
	"log"

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
    // Capture connection properties.
	
    cfg := mysql.Config{
        User:   "root" ,
        Passwd: "root",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
		AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
	// get all albums from database

	allAlbums,err := getAllAlbums()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("all albums: %v\n", allAlbums)

	// get all albums where artist name is John Coltrane
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
	// Hard-code ID 2 here to test the query.
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
	// the album befor updating
	alb, err = albumByID(albID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album befor updating: %v\n", alb)

	/* update the added row */
	albID, err = updateAlbum(Album{
		Title:  "HABIBI",
		Artist: "Faouzia",
		Price:  39,
	},albID)
	if err != nil {
		log.Fatal(err)
	}


	// the album after updating
	alb, err = albumByID(albID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("updated album: %v\n", alb)

	/* delete the updated album*/

	err = deleteAlbum(albID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of deleted album: %v\n", albID)
	allAlbums,err = getAllAlbums()
	if err != nil {
		log.Fatal(err)
	}
	/* print all albums */
	fmt.Printf("all albums: %v\n", allAlbums)
	
}
func getAllAlbums() ([]Album, error)   {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album")

	if err != nil {
		return nil, fmt.Errorf("albums : %v",  err)
	}

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
            return nil, fmt.Errorf("albums : %v", err)
        }
		albums = append(albums, album)
	}
	return albums, nil
}
// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album

    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}

func albumByID(id int64) (Album, error) {
    // An album to hold data from the returned row.
    var alb Album

    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumsById %d: no such album", id)
        }
        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}

func deleteAlbum(id int64) error {
	// An album to hold data from the returned row.
    

    _, err := db.Exec("DELETE FROM album WHERE id = ?", id)
	
    if err != nil {
        return  fmt.Errorf("addAlbum: %v", err)
    }
    return  nil
}

func updateAlbum(alb Album, id int64) (int64, error) {
	
    _, err := db.Exec("UPDATE album set title = ?, artist = ?, price = ? WHERE id = ?", alb.Title, alb.Artist, alb.Price,id)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    
    return id, nil
}