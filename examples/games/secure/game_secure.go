package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	obs "github.com/Dentrax/obscure-go/observer"
	secure "github.com/Dentrax/obscure-go/types"
)

func main() {
	w := obs.CreateWatcher("watcher")

	score := secure.NewInt(0)
	coin := secure.NewInt(0)

	score.AddWatcher(w)
	coin.AddWatcher(w)

	scanner := bufio.NewScanner(os.Stdin)

	println("Simple Game - [i: inc, d: dec, q: quit]")
	println("Objective: Reach 1000 and win!")
	println("Keys: [i: inc, d: dec, q: quit]")

	timeStart := time.Now()

	var visited = make(map[int]bool)

	for {
		println()
		print("Type: ")

		scanner.Scan()

		switch scanner.Text() {
		case "i":
			score.Inc()
			break
		case "d":
			if score.Get() > 0 {
				score.Dec()
			}
			break
		case "q":
			os.Exit(0)
		default:
			println("Allowed keys: [i: inc, d: dec, q: quit]")
			continue
		}

		println("Score: ", score.Get())

		if score.Get() % 10 == 0 {
			if _, ok := visited[score.Get()]; !ok {
				visited[score.Get()] = true
				coin.Inc()
				println("Coin: ", coin.Get())
			}
		}

		if score.Get() >= 100 {
			timeElapsed := time.Since(timeStart)
			println()
			println(fmt.Sprintf("You won! Score: %d, Coin: %d, Time: %s", score.Get(), coin.Get(), timeElapsed))
			os.Exit(0)
		}
	}
}
