package usecase

import (
	"fmt"
	"time"

	"number-guessing-game/internal/entity"
	"number-guessing-game/internal/repository"
)

// GameUsecase contains the business logic for the game
type GameUsecase struct {
	randomRepo repository.RandomRepository
}

// NewGameUsecase creates a new instance of GameUsecase
func NewGameUsecase(randomRepo repository.RandomRepository) *GameUsecase {
	return &GameUsecase{randomRepo: randomRepo}
}

// StartGame initializes a new game based on the selected difficulty
func (uc *GameUsecase) StartGame(difficulty string) (*entity.Game, error) {
	attempts := uc.getAttemptsByDifficulty(difficulty)
	if attempts == 0 {
		return nil, fmt.Errorf("invalid difficulty level")
	}

	target := uc.randomRepo.GenerateRandomNumber(1, 100)
	return &entity.Game{
		TargetNumber: target,
		AttemptsLeft: attempts,
		StartTime:    time.Now(),
	}, nil
}

// MakeGuess checks the user's guess and returns feedback
func (uc *GameUsecase) MakeGuess(game *entity.Game, guess int) (string, bool) {
	game.AttemptsLeft--

	if guess == game.TargetNumber {
		return "Congratulations! You guessed the correct number.", true
	}

	if guess < game.TargetNumber {
		return "Incorrect! The number is greater than your guess.", false
	}

	return "Incorrect! The number is less than your guess.", false
}

// getAttemptsByDifficulty determines the number of attempts based on difficulty
func (uc *GameUsecase) getAttemptsByDifficulty(difficulty string) int {
	switch difficulty {
	case "easy":
		return 10
	case "medium":
		return 5
	case "hard":
		return 3
	default:
		return 0
	}
}
