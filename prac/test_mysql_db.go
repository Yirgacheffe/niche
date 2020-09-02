package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("user:password@tcp(127.0.0.1:3306)/hello-world")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// check database right away
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// query by id
	var (
		id   int
		name string
	)

	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// query by statement
	/*
		stmt, err := db.Prepare("select id, name from users where id = ?")
		if err != nil {
			log.Fatal(err)
		}

		defer stmt.Close()

		rows, err = stmt.Query(1)
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
		}

		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
	*/

	/*
		var name string
		err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(name)
	*/

	/*
		stmt, err := db.Prepare("select name from users where id = ?")
		if err != nil {
			log.Fatal(err)
		}

		defer stmt.Close()

		var name string
		err = stmt.QueryRow(1).Scan(&name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(name)
	*/

	// insert new record
	stmt, err := db.Prepare("insert into users(name) values(?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec("Dolly")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	// statement in tx
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	stmt, err = tx.Prepare("insert into foo values(?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// stmt.Close() runs here!

	// handle errors
	/*
		for rows.Next() {
			// ...
		}

		if err = rows.Err(); err != nil {
			// handle err here
		}

		for rows.Next() {
			break
		}

		if err = rows.Close(); err != nil {
			log.Print(err)
		}

		var name string
		err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
		if err != nil {
			if err == sql.ErrNoRows {
				// No rows, means no error
			} else {
				log.Fatal(err)
			}
		}

		fmt.Println(name)
	*/

	/*
		rows, err := db.Query("select * from sometalbe")
		if strings.Contains(err.Error(), "Access denied") {
			// Handle error here
		}
	*/

}
