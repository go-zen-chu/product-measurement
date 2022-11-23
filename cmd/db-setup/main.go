package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:rootpassword@(localhost:3306)/local")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = db.Close(); err != nil {
			panic(err)
		}
	}()
	b, err := os.ReadFile("./db/schema/jira_tables.sql")
	if err != nil {
		panic(err)
	}
	sqldump := string(b)
	query := ""
	for _, line := range strings.Split(sqldump, "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, ";") {
			spls := strings.Split(line, ";")
			for _, spl := range spls[:len(spls)-1] {
				query += spl + ";"
				fmt.Printf("%s\n", query)
				if _, err := db.Exec(query); err != nil {
					panic(err)
				}
				query = ""
			}
			query += spls[len(spls)-1]
			continue
		}
		// sql continues to next line
		query += line
	}
	fmt.Println("success")
}
