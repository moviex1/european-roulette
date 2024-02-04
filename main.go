package main

import "example.com/roulette"

func main() {
	er := roulette.BuildEuropeanRoulette()
	er.StartGame()
}
