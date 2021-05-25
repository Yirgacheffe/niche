package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"math"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello-world")
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
	stmt, err := db.Prepare("select id, name from users where id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows1, err = stmt.Query(1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows1.Close()
	for rows1.Next() {
		err = rows1.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = rows1.Err(); err != nil {
		log.Fatal(err)
	}
	
	// Query row
	var name1 string
	err = db.QueryRow("select name from users where id = ?", 1).Scan(&name1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name1)
	

	// Query by prepared
	stmt1, err := db.Prepare("select name from users where id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt1.Close()

	var name2 string
	err = stmt.QueryRow(1).Scan(&name2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name2)
	
	// insert new record
	stmt2, err := db.Prepare("insert into users(name) values(?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt2.Exec("Dolly")
	defer stmt2.Close()
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
	stmt3, err = tx.Prepare("insert into foo values(?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt3.Close()
	for i := 0; i < 10; i++ {
		_, err = stmt3.Exec(i)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	stmt3.Close() runs here!

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
	*/

	var namex string
	err = db.QueryRow("select name from users where id = ?", 1).Scan(&namex)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows, means no error
			fmt.Println("No name found")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println(namex)

	rows, err = db.Query("select * from sometalbe")
	if strings.Contains(err.Error(), "Access denied") {
		// Handle error here
	}

	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == 1045 {
			// Handle the permission-denied error
		}

		// VividCortex err number
		// if driverErr.Number == MySQLError.ER_ACCESS_DENIED_ERROR {
		//
		// }
	}

	// Null type
	for rows.Next() {
		var s sql.NullString
		err := rows.Scan(&s)

		if s.Valid {
			// normal case
		} else {
			// Handle null
		}
	}

	// Null value
	rowss, err = db.Query(`select name, coalesce(other_field,'') as otherField where id = ?`, 43)
	for rowss.Next() {
		err = rowss.Scan(&name, &otherField)
		// ..
		// If `other_field` was NULL, `otherField` is now an empty string.
	}

	// don't know the column
	cols, err := rowss.Columns()
	if err != nil {
		// handle the error
	} else {
		dest := []interface{}{
			new(uint64), // id
			new(string), // host
			new(string), // user
			new(string), // db
			new(string), // command
			new(uint32), // time
			new(string), // state
			new(string), // info
		}

		if len(cols) == 11 {
			// per server
		} else if len(cols) > 8 {
			// handle this case
		}

		err = rowss.Scan(dest...)
	}

	// RAW types
	vals := make([]interface, len(cols))
	for i, _ := range cols {
		vals[i] = new(sql.RawBytes)
	}

	for rowsd.Next() {
		err = rowsd.Scan(vals...)
	}

	// uint64 not support
	_, err = db.Exec("insert into user(id) values(?)", math.MaxUint64)

}
