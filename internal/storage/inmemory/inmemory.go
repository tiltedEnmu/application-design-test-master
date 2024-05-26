package inmemory

import (
	"sync"

	"applicationDesignTest/internal/domain/models"
)

type Storage struct {
	mu    *sync.RWMutex
	rooms map[string][]models.BookedDay
}

func New() *Storage {
	return &Storage{
		mu:    &sync.RWMutex{},
		rooms: make(map[string][]models.BookedDay),
	}
}
