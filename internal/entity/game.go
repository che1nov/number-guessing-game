package entity

import "time"

// Game represents the state of the game
type Game struct {
	TargetNumber int       // The randomly generated number to guess
	AttemptsLeft int       // Number of attempts remaining
	StartTime    time.Time // Time when the game started
}
