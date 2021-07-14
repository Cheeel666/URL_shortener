package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	sq "github.com/Masterminds/squirrel"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = ""
	db_name = "test"
)



func AddUrl(url string) sq.InsertBuilder {
	query := sq.Insert("links").Columns("url").Values(url)
	fmt.Println(query.ToSql())
	//rows, err := query.RunWith
	return query
}

func UpdateUrl(id int, shortUrl string)sq.UpdateBuilder {
	query := sq.Update("links").Where("id=")
	return query
}

func main(){
	pg_con_string := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name)
	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(db.QueryRow("insert into links (url) values ($1)", 123))

	// We can also ping our connection which will let us know if our connection is correct /// or not then we put an error-handling code right after that.

	//query := sq.Select("*").From("link")
	//
	//sql, args, err := query.ToSql()
	//fmt.Println(sql, args, err)
}