package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func runSQLDump(db *sql.DB, path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	b, err := os.ReadFile(path)
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
	return nil
}

func setupTables(db *sql.DB) error {
	if err := runSQLDump(db, "./db/schema/jira_tables.sql"); err != nil {
		return fmt.Errorf("setup jira table: %w", err)
	}
	return nil
}

func main() {
	db, err := sql.Open("mysql", "root:rootpassword@(localhost:3306)/")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = db.Close(); err != nil {
			panic(err)
		}
	}()
	if err := runSQLDump(db, "./db/schema/databases.sql"); err != nil {
		panic(fmt.Errorf("setup database: %w", err))
	}
	// setup database for production
	pmdb, err := sql.Open("mysql", "root:rootpassword@(localhost:3306)/PM")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = pmdb.Close(); err != nil {
			panic(err)
		}
	}()
	if err := setupTables(pmdb); err != nil {
		panic(err)
	}
	// setup database for tests
	pmtestdb, err := sql.Open("mysql", "root:rootpassword@(localhost:3306)/PM_TEST")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = pmtestdb.Close(); err != nil {
			panic(err)
		}
	}()
	if err := setupTables(pmtestdb); err != nil {
		panic(err)
	}
	fmt.Println("db setup success")
}
