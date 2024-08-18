package routes

import (
	"github.com/0mwa/testtask-go-url-shortener/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type Resolver struct {
	storage storage.Storage
}

func NewResolver(storage storage.Storage) *Resolver {
	return &Resolver{
		storage: storage,
	}
}

func (r *Resolver) ResolveURL(c *fiber.Ctx) error {
	shortURL := c.Params("url")

	originalURL, err := r.storage.Read(shortURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Redirect(originalURL, fiber.StatusMovedPermanently)
}
