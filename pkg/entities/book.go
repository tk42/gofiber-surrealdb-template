package entities

import (
	"time"
)

// Book Constructs your Book model under entities.
type Book struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// DeleteRequest struct is used to parse Delete Requests for Books
type DeleteRequest struct {
	ID string `json:"id"`
}
