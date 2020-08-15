package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var score uint64 = 0
	var coin uint64 = 0

	scanner := bufio.NewScanner(os.Stdin)

	println("Simple Game - [i: inc, d: dec, q: quit]")
	println("Objective: Reach 1000 and win!")
	println("Keys: [i: inc, d: dec, q: quit]")

	timeStart := time.Now()

	var visited = make(map[uint64]bool)

	for {
		println()
		print("Type: ")

		scanner.Scan()

		switch scanner.Text() {
		case "i":
			score += 1
			break
		case "d":
			if score > 0 {
				score -= 1
			}
			break
		case "q":
			os.Exit(0)
		default:
			println("Allowed keys: [i: inc, d: dec, q: quit]")
			continue
		}

		println("Score: ", score)

		if score % 10 == 0 {
			if _, ok := visited[score]; !ok {
				visited[score] = true
				coin += 1
				println("Coin: ", coin)
			}
		}

		if score >= 100 {
			timeElapsed := time.Since(timeStart)
			println()
			println(fmt.Sprintf("You won! Score: %d, Coin: %d, Time: %s", score, coin, timeElapsed))
			os.Exit(0)
		}
	}
}
