# Number Guessing Game

Welcome to the Number Guessing Game! Test your luck and see if you can guess the number that the computer randomly selects.

## How to Play

This is a command-line interface (CLI) based game. Follow the instructions below to play the game:

1. When the game starts, a welcome message along with the rules of the game will be displayed.
2. The computer will randomly select a number between 1 and 100.
3. You will be prompted to select a difficulty level which will determine the number of chances you get to guess the number:
    - Easy: 10 chances
    - Medium: 5 chances
    - Hard: 3 chances
4. Enter your guess.
5. If your guess is correct, a congratulatory message will be displayed along with the number of attempts it took to guess the number.
6. If your guess is incorrect, a message will be displayed indicating whether the number is greater or less than your guess.
7. The game will end when you guess the correct number or run out of chances.

## Sample Output

```
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

Enter your guess: 50
Incorrect! The number is less than 50.

Enter your guess: 25
Incorrect! The number is greater than 25.

Enter your guess: 35
Incorrect! The number is less than 35.

Enter your guess: 30
Congratulations! You guessed the correct number in 4 attempts.
```

## Additional Features

To make the game more interesting, the following features are included:

- **Multiple Rounds**: After each round, you will be asked if you want to play again.
- **Timer**: A timer will keep track of how long it takes you to guess the number.
- **Hint System**: If you are stuck, the game can provide you with hints.
- **High Score**: The game will keep track of your high score, i.e., the fewest number of attempts it took to guess the number under a specific difficulty level.

## Getting Started

### Prerequisites

- Go 1.15 or later

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your_username/number-guessing-game-go.git
   ```
2. Navigate to the project directory:
   ```sh
   cd number-guessing-game-go
   ```

### Running the Game

Run the following command to start the game:
```sh
go run main.go
```

### Building the Game (Optional)

If you want to build the game into an executable, run:
```sh
go build -o number-guessing-game main.go
./number-guessing-game
```

### Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Feedback

Submit your solution and get feedback from the community. Upvote if you find the game enjoyable!

Happy guessing!

https://roadmap.sh/projects/number-guessing-game