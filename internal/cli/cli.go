package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"number-guessing-game/internal/entity"
	"number-guessing-game/internal/usecase"
)

// CLI manages the command-line interface
type CLI struct {
	usecase *usecase.GameUsecase
}

// NewCLI creates a new instance of CLI
func NewCLI(usecase *usecase.GameUsecase) *CLI {
	return &CLI{usecase: usecase}
}

// Run starts the CLI
func (c *CLI) Run() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You will have a limited number of chances to guess the correct number.\n")

	for {
		difficulty := selectDifficulty()
		game, err := c.usecase.StartGame(difficulty)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Great! Let's start the game!\n\n")

		won := playGame(c.usecase, game)
		if won {
			duration := time.Since(game.StartTime).Seconds()
			fmt.Printf("You guessed the correct number in %.2f seconds.\n", duration)
		} else {
			fmt.Printf("Sorry, you've run out of attempts. The correct number was %d.\n", game.TargetNumber)
		}

		if !askToPlayAgain() {
			fmt.Println("Thank you for playing! Goodbye!")
			break
		}
	}
}

// selectDifficulty prompts the user to select a difficulty level
func selectDifficulty() string {
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	for {
		var choice string
		fmt.Print("Enter your choice (1/2/3): ")
		fmt.Scanln(&choice)

		switch strings.TrimSpace(choice) {
		case "1":
			return "easy"
		case "2":
			return "medium"
		case "3":
			return "hard"
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}
	}
}

// playGame manages the guessing process
func playGame(usecase *usecase.GameUsecase, game *entity.Game) bool {
	for game.AttemptsLeft > 0 {
		guess := getUserGuess(game.AttemptsLeft)
		message, won := usecase.MakeGuess(game, guess)
		fmt.Println(message)
		if won {
			return true
		}
	}
	return false
}

// getUserGuess gets the user's guess
func getUserGuess(attemptsLeft int) int {
	for {
		var input string
		fmt.Printf("Attempt %d: Enter your guess (1-100): ", attemptsLeft)
		fmt.Scanln(&input)

		guess, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || guess < 1 || guess > 100 {
			fmt.Println("Invalid input. Please enter a number between 1 and 100.")
			continue
		}
		return guess
	}
}

// askToPlayAgain asks if the user wants to play again
func askToPlayAgain() bool {
	for {
		var response string
		fmt.Print("Would you like to play again? (yes/no): ")
		fmt.Scanln(&response)

		response = strings.ToLower(strings.TrimSpace(response))
		if response == "yes" || response == "y" {
			return true
		} else if response == "no" || response == "n" {
			return false
		} else {
			fmt.Println("Invalid response. Please enter 'yes' or 'no'.")
		}
	}
}
