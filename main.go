package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Game struct {
	targetValue     int
	numberOfGuesses int
	rightGuess      bool
}

func newGame(maxVal int) (*Game, error) {
	if maxVal < 0 {
		return nil, errors.New("Max value can't be negative")
	}

	return &Game{
		targetValue:     rand.Intn(maxVal),
		numberOfGuesses: 3,
		rightGuess:      false,
	}, nil
}

func (g *Game) Start() {
	var guessingVal int
	currentGuess := 0

	for ; currentGuess < g.numberOfGuesses-1; currentGuess++ {
		fmt.Println("Guess the number: ")
		fmt.Scanln(&guessingVal)

		if guessingVal < g.targetValue {
			fmt.Println("The Number is larger")
		} else if guessingVal > g.targetValue {
			fmt.Println("The number is smaller")
		} else {
			g.rightGuess = true
			break
		}
	}

	if g.rightGuess {
		fmt.Printf("You won\nAmount of Guesses: %d\n", currentGuess+1)
	} else {
		fmt.Printf("Game Over\nThe right number was: %d\n", g.targetValue)
	}
}

func main() {
	game, err := newGame(20)
	if err != nil {
		fmt.Println(err)
		return
	}

	game.Start()
}
