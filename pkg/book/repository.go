package book

import (
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/tk42/gofiber-surrealdb-template/pkg/entities"

	surrealdb "github.com/surrealdb/surrealdb.go"
)

// Repository interface allows us to access the CRUD Operations in surrealdb here.
type Repository interface {
	CreateBook(book *entities.Book) (*entities.Book, error)
	ReadBook() (*[]entities.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(ID string) error
}
type repository struct {
	Database *surrealdb.DB
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *surrealdb.DB) Repository {
	return &repository{
		Database: collection,
	}
}

// CreateBook is a surrealdb repository that helps to create books
func (r *repository) CreateBook(_book *entities.Book) (*entities.Book, error) {
	_book.CreatedAt = time.Now()
	_book.UpdatedAt = time.Now()
	b, err := r.Database.Create("book", structs.Map(*_book))
	if err != nil {
		return nil, err
	}
	var book entities.Book
	err = surrealdb.Unmarshal(b, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// ReadBook is a surrealdb repository that helps to fetch books
func (r *repository) ReadBook() (*[]entities.Book, error) {
	var books []entities.Book
	cursor, err := r.Database.Select("book")
	if err != nil {
		return nil, err
	}
	err = surrealdb.Unmarshal(cursor, &books)
	if err != nil {
		return nil, err
	}
	return &books, nil
}

// UpdateBook is a surrealdb repository that helps to update books
func (r *repository) UpdateBook(_book *entities.Book) (*entities.Book, error) {
	var books []entities.Book
	cursor, err := r.Database.Select("book")
	if err != nil {
		return nil, err
	}
	err = surrealdb.Unmarshal(cursor, &books)
	if err != nil {
		return nil, err
	}

	for _, book := range books {
		if book.ID == "book:"+_book.ID {
			book.UpdatedAt = time.Now()
			if _book.Author != "" {
				book.Author = _book.Author
			}
			if _book.Title != "" {
				book.Title = _book.Title
			}
			_, err := r.Database.Update("book:"+_book.ID, book)
			if err != nil {
				return nil, err
			}
			return &book, nil
		}
	}
	return nil, fmt.Errorf("BookID(%s) not found", _book.ID)
}

// DeleteBook is a surrealdb repository that helps to delete books
func (r *repository) DeleteBook(ID string) error {
	_, err := r.Database.Delete("book:" + ID)
	if err != nil {
		return err
	}
	return nil
}
