package roulette

import "testing"

var bet = 10.0

type RouletteNumberTests struct {
	want           float64
	rouletteNumber RouletteNumber
}

var ruleTests = []struct {
	name                string
	rule                RouletteRule
	rouletteNumberTests []RouletteNumberTests
}{
	{name: Even, rule: &EvenRule{Even, bet}, rouletteNumberTests: []RouletteNumberTests{
		{20.0, RouletteNumber{num: 10, color: "black"}},
		{0.0, RouletteNumber{num: 9, color: "red"}},
	}},
	{name: Odd, rule: &OddRule{Odd, bet}, rouletteNumberTests: []RouletteNumberTests{
		{0.0, RouletteNumber{num: 10, color: "black"}},
		{20.0, RouletteNumber{num: 9, color: "red"}},
	}},
	{name: FirstEighteen, rule: &FirstEighteenRule{FirstEighteen, bet}, rouletteNumberTests: []RouletteNumberTests{
		{20.0, RouletteNumber{num: 10, color: "black"}},
		{0.0, RouletteNumber{num: 28, color: "black"}},
	}},
	{name: SecondEighteen, rule: &SecondEighteenRule{SecondEighteen, bet}, rouletteNumberTests: []RouletteNumberTests{
		{0.0, RouletteNumber{num: 10, color: "black"}},
		{20.0, RouletteNumber{num: 28, color: "black"}},
	}},
	{name: FirstTwelve, rule: &FirstTwelveRule{FirstTwelve, bet}, rouletteNumberTests: []RouletteNumberTests{
		{30.0, RouletteNumber{num: 10, color: "black"}},
		{0.0, RouletteNumber{num: 14, color: "red"}},
		{0.0, RouletteNumber{num: 28, color: "black"}},
	}},
	{name: SecondTwelve, rule: &SecondTwelveRule{SecondTwelve, bet}, rouletteNumberTests: []RouletteNumberTests{
		{0.0, RouletteNumber{num: 10, color: "black"}},
		{30.0, RouletteNumber{num: 14, color: "red"}},
		{0.0, RouletteNumber{num: 28, color: "black"}},
	}},
	{name: ThirdTwelve, rule: &ThirdTwelveRule{ThirdTwelve, bet}, rouletteNumberTests: []RouletteNumberTests{
		{0.0, RouletteNumber{num: 10, color: "black"}},
		{0.0, RouletteNumber{num: 14, color: "red"}},
		{30.0, RouletteNumber{num: 28, color: "black"}},
	}},
	{name: RedColor, rule: &RedColorRule{RedColor, bet}, rouletteNumberTests: []RouletteNumberTests{
		{0.0, RouletteNumber{num: 10, color: "black"}},
		{20.0, RouletteNumber{num: 9, color: "red"}},
	}},
	{name: BlackColor, rule: &BlackColorRule{BlackColor, bet}, rouletteNumberTests: []RouletteNumberTests{
		{20.0, RouletteNumber{num: 10, color: "black"}},
		{0.0, RouletteNumber{num: 9, color: "red"}},
	}},
	{name: ConcreteNumber, rule: &ConcreteNumberRule{ConcreteNumber, bet, 10}, rouletteNumberTests: []RouletteNumberTests{
		{360.0, RouletteNumber{num: 10, color: "black"}},
		{0.0, RouletteNumber{num: 9, color: "red"}},
	}},
}

func TestRules(t *testing.T) {
	for _, tt := range ruleTests {
		t.Run(tt.name, func(t *testing.T) {
			for _, rnt := range tt.rouletteNumberTests {
				got := tt.rule.Eval(rnt.rouletteNumber)
				if rnt.want != got {
					t.Errorf("got %g, want %g", got, rnt.want)
				}
			}
		})
	}
}
