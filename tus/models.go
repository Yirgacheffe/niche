package main

import "time"

// File type means... the file client want to upload to the server
type File struct {
	ID           int64     `json:"id"`
	Offset       int       `json:"offset"`
	UploadLength int       `json:"upload_length"`
	IsComplete   string    `json:"is_complete"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
