package main

import (
	"fmt"
	"mooc/db"
)

func main() {
	// columns := make([]db.WebSite, 0)
	// columns = db.Query("select * from website")
	// fmt.Println(columns)
	affect := db.Insert("insert into website(name,title,comment,href) values(?,?,?,?)", "dragon", "dra", "test", "ahref")
	fmt.Println(affect)
	// update := db.Insert("update website set title =? where id=?", "update test", 1)
	// fmt.Println(update)
	// del := db.Delete("delete from website where id=?", 2)
	// fmt.Println(del)
}
