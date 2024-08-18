package storage

type Storage interface {
	Write(shortURL, originalURL string) error
	Read(shortURL string) (string, error)
}
