package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = ""
	db_name = "postgres"
)



func AddUrl(url string) int {
	var id int
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows := db.QueryRow("INSERT INTO links (url) VALUES ($1)", url)
	fmt.Println(rows)
	row, err := db.Query("SELECT max(id_url) from links")
	if err != nil {
		panic(err)
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertNew: ", id, url)
	return id
}

func UpdateUrl(id int, shortUrl string) {
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Update url:", id, shortUrl)

	rows := db.QueryRow("UPDATE links SET shorten_url = ($1) where id_url = ($2);", shortUrl, id)
	fmt.Println(rows)
}

func SelectLongUrl(url string) string {
	var resultUrl string
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query("SELECT url FROM links where shorten_url = ($1)", url)
	for row.Next() {
		if err := row.Scan(&resultUrl); err != nil {
			log.Fatal(err)
		}
	}
	return resultUrl
}

//
//func main(){
//
//	//query := sq.Select("*").From("links")
//	//
//	//sql, args, err := query.ToSql()
//	//fmt.Println(sql, args, err)
//}