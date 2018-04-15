package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type WebSite struct {
	Id      int
	Name    string
	Title   string
	Comment string
	Href    string
}

var mydb *sql.DB

func newCreateDB() *sql.DB {
	// handle := "ranrl:password@tcp(localhost:3306)/test?charset=utf8"
	handle := user + ":" + pwd + "@" + proto + "(" + host + ":" + port + ")/" + dbname + "?charset=" + encode
	if mydb == nil {
		var err error
		mydb, err = sql.Open("mysql", handle)
		checkErr(err)
	}
	return mydb
}
func Query(selectSql string) []WebSite {
	fmt.Println(selectSql)
	db := newCreateDB()
	rows, err := db.Query(selectSql)
	checkErr(err)
	defer rows.Close()
	// columns, err := rows.Columns()
	// fmt.Println(columns)
	// checkErr(err)
	data := make([]WebSite, 0)
	for rows.Next() {
		var id int
		var name string
		var title string
		var comment string
		var href string
		if err := rows.Scan(&id, &name, &title, &comment, &href); err != nil {
			checkErr(err)
		}
		fmt.Println(id, name, title, comment, href)
		var web WebSite
		web.Id = id
		web.Title = title
		web.Comment = comment
		web.Href = href
		data = append(data, web)
	}
	return data
}

func Insert(insertSql string, args ...interface{}) int64 {
	db := newCreateDB()
	stmt, err := db.Prepare(insertSql)
	checkErr(err)
	result, err := stmt.Exec(args...)
	checkErr(err)
	affect, err := result.RowsAffected()
	checkErr(err)
	return affect
}

func Update(updateSql string, args ...interface{}) int64 {
	db := newCreateDB()
	stmt, err := db.Exec(updateSql, args...)
	checkErr(err)
	update, err := stmt.RowsAffected()
	checkErr(err)
	return update
}

func Delete(deleteSql string, args ...interface{}) int64 {
	db := newCreateDB()
	result, err := db.Exec(deleteSql, args...)
	checkErr(err)
	delNum, err := result.RowsAffected()
	checkErr(err)
	return delNum
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
