package roulette

import "fmt"

type RouletteTable struct {
	rules []RouletteRule
}

func (t *RouletteTable) AddRule(r RouletteRule) {
	t.rules = append(t.rules, r)
}

func (t *RouletteTable) CalculateWin(rn RouletteNumber) float64 {
	win := .0
	for _, rule := range t.rules {
		win += rule.Eval(rn)
	}

	return win
}

func (t *RouletteTable) ClearTable() {
	t.rules = nil
}

func (t *RouletteTable) DisplayBetsResult(rn RouletteNumber) {
	for _, rule := range t.rules {
		fmt.Printf("%s - %g \n", rule.GetName(), rule.Eval(rn))
	}
}

func BuildRouletteTable() RouletteTable {
	return RouletteTable{}
}
