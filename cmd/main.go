package main

import (
	"flag"
	"github.com/0mwa/testtask-go-url-shortener/internal/database"
	"github.com/0mwa/testtask-go-url-shortener/internal/routes"
	"github.com/0mwa/testtask-go-url-shortener/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

func setupStorage() storage.Storage {
	useDb := flag.Bool("d", false, "if true use db storage, else use memory storage")
	flag.Parse()
	if *useDb {
		db, err := database.NewPostgresClient()
		if err != nil {
			panic(err)
		}
		return storage.NewPgStorage(db)
	} else {
		return storage.NewMemoryStorage()
	}
}

func setupRoutes(app *fiber.App) {

	storage_ := setupStorage()

	s := routes.NewShortener(storage_)
	r := routes.NewResolver(storage_)

	app.Get("/:url", r.ResolveURL)
	app.Post("/", s.ShortenURL)
}
