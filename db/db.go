package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	db_name = "postgres"
)



func AddUrl(url string) (id int, err error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	rows := db.QueryRow("INSERT INTO links (url) VALUES ($1)", url)
	fmt.Println(rows)
	row, err := db.Query("SELECT max(id_url) from links")
	if err != nil {
		return -1, err
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}

	//fmt.Println("InsertNew: ", id, url)
	return id, nil
}

func UpdateUrl(id int, shortUrl string) (err error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	rows := db.QueryRow("UPDATE links SET shorten_url = ($1) where id_url = ($2);", shortUrl, id)
	_ = rows
	//fmt.Println(rows)
	return nil
}

func SelectLongUrl(url string) (resultUrl string, err error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return "", err
	}
	defer db.Close()

	row, err := db.Query("SELECT url FROM links where shorten_url = ($1)", url)
	for row.Next() {
		if err := row.Scan(&resultUrl); err != nil {
			return "", err
		}
	}
	return resultUrl, nil
}

//
//func main(){
//
//	//query := sq.Select("*").From("links")
//	//
//	//sql, args, err := query.ToSql()
//	//fmt.Println(sql, args, err)
//}