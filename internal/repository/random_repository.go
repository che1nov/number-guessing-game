package repository

import (
	"math/rand"
	"time"
)

// RandomRepository defines an interface for generating random numbers
type RandomRepository interface {
	GenerateRandomNumber(min, max int) int
}

// DefaultRandomRepository implements the RandomRepository interface
type DefaultRandomRepository struct{}

// NewDefaultRandomRepository creates a new instance of DefaultRandomRepository
func NewDefaultRandomRepository() *DefaultRandomRepository {
	return &DefaultRandomRepository{}
}

// GenerateRandomNumber generates a random number between min and max (inclusive)
func (r *DefaultRandomRepository) GenerateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
