package routes

import (
	"fmt"
	"github.com/0mwa/testtask-go-url-shortener/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type Resolver struct {
	storage storage.Storage
}
type resolverResponse struct {
	OriginalURL string `json:"url"`
}

func NewResolver(storage storage.Storage) *Resolver {
	return &Resolver{
		storage: storage,
	}
}

func (r *Resolver) ResolveURL(c *fiber.Ctx) error {
	shortURL := c.Params("url")

	response_, err := r.createResponse(shortURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response_)
}

func (r *Resolver) createResponse(shortURL string) (resolverResponse, error) {
	originalURL, err := r.storage.Read(shortURL)
	if err != nil {
		return resolverResponse{}, fmt.Errorf(err.Error())
	}

	return resolverResponse{
		OriginalURL: originalURL,
	}, nil
}
