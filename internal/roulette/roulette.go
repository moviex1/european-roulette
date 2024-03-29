package roulette

import "math/rand"

var rouletteNumbers = []RouletteNumber{
	{num: 0, color: "green"},
	{num: 1, color: "red"},
	{num: 2, color: "black"},
	{num: 3, color: "red"},
	{num: 4, color: "black"},
	{num: 5, color: "red"},
	{num: 6, color: "black"},
	{num: 7, color: "red"},
	{num: 8, color: "black"},
	{num: 9, color: "red"},
	{num: 10, color: "black"},
	{num: 11, color: "black"},
	{num: 12, color: "red"},
	{num: 13, color: "black"},
	{num: 14, color: "red"},
	{num: 15, color: "black"},
	{num: 16, color: "red"},
	{num: 17, color: "black"},
	{num: 18, color: "red"},
	{num: 19, color: "red"},
	{num: 20, color: "black"},
	{num: 21, color: "red"},
	{num: 22, color: "black"},
	{num: 23, color: "red"},
	{num: 24, color: "black"},
	{num: 25, color: "red"},
	{num: 26, color: "black"},
	{num: 27, color: "red"},
	{num: 28, color: "black"},
	{num: 29, color: "black"},
	{num: 30, color: "red"},
	{num: 31, color: "black"},
	{num: 32, color: "red"},
	{num: 33, color: "black"},
	{num: 34, color: "red"},
	{num: 35, color: "black"},
	{num: 36, color: "red"},
}

type RouletteNumber struct {
	num   int
	color string
}

type Roulette struct {
	numbers []RouletteNumber
}

func (r *Roulette) Play() RouletteNumber {
	return r.numbers[rand.Int()%37]
}

func BuildRoulette() Roulette {
	roulette := Roulette{
		numbers: rouletteNumbers,
	}

	return roulette
}
