package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tk42/gofiber-surrealdb-template/api/routes"
	"github.com/tk42/gofiber-surrealdb-template/pkg/book"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	surrealdb "github.com/surrealdb/surrealdb.go"
)

var (
	SURREALDB_URL  = os.Getenv("SURREALDB_URL")
	SURREALDB_USER = os.Getenv("SURREALDB_USER")
	SURREALDB_PASS = os.Getenv("SURREALDB_PASS")
)

func main() {
	db, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	bookRepo := book.NewRepo(db)
	bookService := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture surrealdb shop!"))
	})
	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*surrealdb.DB, error) {
	db, err := surrealdb.New(SURREALDB_URL)
	if err != nil {
		return nil, err
	}

	_, err = db.Signin(map[string]interface{}{
		"user": SURREALDB_USER,
		"pass": SURREALDB_PASS,
	})
	if err != nil {
		return nil, err
	}

	// specify Namespace and Database
	_, err = db.Use("ns", "books")
	return db, err
}
