package main

import (
	"fmt"
	"strings"

	"github.com/tzfe/board"
)

func main() {
	// initialize an empty board
	b := board.Board{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	// spawn two initial tile
	b.Spawn()
	b.Spawn()

	steps := 0
	for {
		fmt.Println(b)
		var move string
		fmt.Scanln(&move)

		switch strings.ToLower(move) {
		case "u":
			b.Up()
		case "d":
			b.Down()
		case "l":
			b.Left()
		case "r":
			b.Right()
		default:
			fmt.Println("Invalid move! Enter one of u/d/l/r swipes.")
			continue
		}
		// spawn a new tile randomly
		b.Spawn()
		steps += 1

		gameOver, playerWon := b.GameOver()
		if gameOver {
			if playerWon {
				fmt.Println(fmt.Sprintf("You won in %d steps!", steps))
			} else {
				fmt.Println(fmt.Sprintf("Game over after %d steps. Better luck next time!", steps))
			}
			break
		}
	}
}
