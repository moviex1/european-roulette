package roulette

import "fmt"

const (
	Even           = "Even"
	Odd            = "Odd"
	FirstEighteen  = "First eighteen"
	SecondEighteen = "Second eighteen"
	FirstTwelve    = "First twelve"
	SecondTwelve   = "Second twelve"
	ThirdTwelve    = "Third twelve"
	RedColor       = "Red color"
	BlackColor     = "Black color"
	ConcreteNumber = "Concrete number"
)

type RuleBuilder struct {
	name string
	bet  float64
	num  int
}

func NewRuleBuilder() *RuleBuilder {
	return &RuleBuilder{}
}

func (rb *RuleBuilder) SetName(name string) *RuleBuilder {
	rb.name = name
	return rb
}

func (rb *RuleBuilder) SetBet(bet float64) *RuleBuilder {
	rb.bet = bet
	return rb
}

func (rb *RuleBuilder) SetNum(num int) *RuleBuilder {
	rb.num = num
	return rb
}

func (rb *RuleBuilder) Build() RouletteRule {
	switch rb.name {
	case Even:
		return &EvenRule{rb.name, rb.bet}
	case Odd:
		return &OddRule{rb.name, rb.bet}
	case FirstEighteen:
		return &FirstEighteenRule{rb.name, rb.bet}
	case SecondEighteen:
		return &SecondEighteenRule{rb.name, rb.bet}
	case FirstTwelve:
		return &FirstTwelveRule{rb.name, rb.bet}
	case SecondTwelve:
		return &SecondTwelveRule{rb.name, rb.bet}
	case ThirdTwelve:
		return &ThirdTwelveRule{rb.name, rb.bet}
	case RedColor:
		return &RedColorRule{rb.name, rb.bet}
	case BlackColor:
		return &BlackColorRule{rb.name, rb.bet}
	case ConcreteNumber:
		return &ConcreteNumberRule{fmt.Sprintf("Number %d", rb.num), rb.bet, rb.num}
	}
	return nil
}
