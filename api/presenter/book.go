package presenter

import (
	"time"

	"github.com/tk42/gofiber-surrealdb-template/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

// Book is the presenter object which will be passed in the response by Handler
type Book struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BookSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func BookSuccessResponse(data *entities.Book) *fiber.Map {
	book := Book{
		ID:        data.ID,
		Title:     data.Title,
		Author:    data.Author,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}

// BooksSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func BooksSuccessResponse(data *[]entities.Book) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
