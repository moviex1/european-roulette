package main

import "github.com/moviex1/european-roulette/internal/roulette"

func main() {
	er := roulette.BuildEuropeanRoulette()
	er.StartGame()
}
