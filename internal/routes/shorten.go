package routes

import (
	"encoding/json"
	"fmt"
	"github.com/0mwa/testtask-go-url-shortener/internal/helpers"
	"github.com/0mwa/testtask-go-url-shortener/internal/storage"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
)

type Shortener struct {
	storage storage.Storage
}

func NewShortener(storage storage.Storage) *Shortener {
	return &Shortener{
		storage: storage,
	}
}

type request struct {
	OriginalURL string `json:"url"`
}

type response struct {
	ShortenURL string `json:"short"`
}

func (s *Shortener) ShortenURL(c *fiber.Ctx) error {
	var req request

	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
	}

	if status, err := s.validateURL(&req); err != nil {
		return c.Status(status).JSON(fiber.Map{"error": err.Error()})
	}

	shortURL := uuid.New().String()[:6]

	if err := s.storage.Write(shortURL, req.OriginalURL); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response_ := s.createResponse(req, shortURL)

	return c.Status(fiber.StatusOK).JSON(response_)
}

func (s *Shortener) validateURL(req *request) (int, error) {
	if !govalidator.IsURL(req.OriginalURL) {
		return fiber.StatusBadRequest, fmt.Errorf("invalid url %s", req.OriginalURL)
	}

	if !helpers.DomainError(req.OriginalURL) {
		return fiber.StatusInternalServerError, fmt.Errorf("domain error")
	}

	req.OriginalURL = helpers.EnforceHTTP(req.OriginalURL)
	return fiber.StatusOK, nil
}

func (s *Shortener) createResponse(req request, id string) response {
	domain := os.Getenv("DOMAIN")
	shortenURL := domain + "/" + id

	return response{
		ShortenURL: shortenURL,
	}
}
