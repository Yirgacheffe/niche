package main

import (
	"database/sql"
	"fmt"
)

// FileRepo - repository service interface
type FileRepo interface {
	GetByID(id int64) (File, error)

	Create(f *File) (int64, error)
	Update(f *File) (int64, error)
}

// NewMySQLFileRepo ... create implementation of FileRepo
func NewMySQLFileRepo(Conn *sql.DB) FileRepo {
	return &mysqlFileRepo{
		Conn: Conn,
	}
}

type mysqlFileRepo struct {
	Conn *sql.DB
}

func (m *mysqlFileRepo) Create(f *File) (int64, error) {

	query := "INSERT INTO TUS_FILES SET OFFSET=?, UPLOAD_LENGTH=?, IS_COMPLETE=?"

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(f.Offset, f.UploadLength, f.IsComplete)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()

}

func (m *mysqlFileRepo) Update(f *File) (int64, error) {

	query := "UPDATE TUS_FILES SET OFFSET=?, IS_COMPLETE=?, UPDATED_AT=NOW() WHERE ID=?"

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(f.Offset, f.IsComplete, f.ID)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()

}

func (m *mysqlFileRepo) GetByID(id int64) (File, error) {

	if id <= 0 {
		return File{}, fmt.Errorf("Invalid ID: %d ", id)
	}

	query := "SELECT ID, OFFSET, UPLOAD_LENGTH, IS_COMPLETE, CREATED_AT, UPDATED_AT " +
		"FROM TUS_FILES WHERE ID=?"

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return File{}, err
	}

	defer stmt.Close()

	f := File{}

	row := stmt.QueryRow(id)
	err = row.Scan(
		&f.ID,
		&f.Offset,
		&f.UploadLength,
		&f.IsComplete, &f.CreatedAt, &f.UpdatedAt)

	if err != nil {
		return File{}, err
	}

	// When you read this code that means you have are not coding carefully
	return f, nil

}
