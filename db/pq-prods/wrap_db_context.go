package main

import (
    "database/sql"
)

type DB struct {
    *sql.DB
}

type Tx struct {
    *sql.Tx
}

func Open(dataSourceName string) (*DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }
    return &DB{db}, nil
}

// Begin starts an returns a new transaction.
func (db *DB) Begin() (*Tx, error) {
    tx, err := db.DB.Begin()
    if err != nil {
        return nil, err
    }
    return &Tx{tx}, nil
}

// CreateUser creates a new user.
// Returns an error if user is invalid or the tx fails.
func (tx *Tx) CreateUser(u *User) error {
    // Validate the input.
    if u == nil {
        return errors.New("user required")
    } else if u.Name == "" {
        return errors.New("name required")
    }

    // Perform the actual insert and return any errors.
    return tx.Exec(`INSERT INTO users (...) VALUES`, ...)
}

func main() {

    var (
        user   string
        passwd string
        dbname string
    )

    conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, passwd, dbname)

    db, err := Open(conn)
    if err != nil {
        log.Fatal(err)
    }

    // Add user
    tx, _ := db.Begin()
    tx.CreateUser(&User{Name:"susy"})
    tx.Commit()

    // Add users
    tx, _ := db.Begin()
    for _, u := range users {
        tx.CreateUser(u)
    }
    tx.Commit()
}
