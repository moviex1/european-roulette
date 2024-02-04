package roulette

type RouletteRule interface {
	GetName() string
	Eval(rn RouletteNumber) float64
}

type EvenRule struct {
	name string
	bet  float64
}

func (r *EvenRule) GetName() string {
	return r.name
}

func (r *EvenRule) Eval(rn RouletteNumber) float64 {
	if rn.num != 0 && (rn.num%2) == 0 {
		return r.bet * 2
	}

	return 0
}

type OddRule struct {
	name string
	bet  float64
}

func (r *OddRule) GetName() string {
	return r.name
}

func (r *OddRule) Eval(rn RouletteNumber) float64 {
	if (rn.num % 2) == 1 {
		return r.bet * 2
	}

	return 0
}

type FirstEighteenRule struct {
	name string
	bet  float64
}

func (r *FirstEighteenRule) GetName() string {
	return r.name
}

func (r *FirstEighteenRule) Eval(rn RouletteNumber) float64 {
	if rn.num > 0 && rn.num < 19 {
		return r.bet * 2
	}

	return 0
}

type SecondEighteenRule struct {
	name string
	bet  float64
}

func (r *SecondEighteenRule) GetName() string {
	return r.name
}

func (r *SecondEighteenRule) Eval(rn RouletteNumber) float64 {
	if rn.num > 18 && rn.num < 37 {
		return r.bet * 2
	}

	return 0
}

type RedColorRule struct {
	name string
	bet  float64
}

func (r *RedColorRule) GetName() string {
	return r.name
}

func (r *RedColorRule) Eval(rn RouletteNumber) float64 {
	if rn.color == "red" {
		return r.bet * 2
	}

	return 0
}

type BlackColorRule struct {
	name string
	bet  float64
}

func (r *BlackColorRule) GetName() string {
	return r.name
}

func (r *BlackColorRule) Eval(rn RouletteNumber) float64 {
	if rn.color == "black" {
		return r.bet * 2
	}

	return 0
}

type FirstTwelveRule struct {
	name string
	bet  float64
}

func (r *FirstTwelveRule) GetName() string {
	return r.name
}

func (r *FirstTwelveRule) Eval(rn RouletteNumber) float64 {
	if rn.num > 0 && rn.num < 13 {
		return r.bet * 3
	}

	return 0
}

type SecondTwelveRule struct {
	name string
	bet  float64
}

func (r *SecondTwelveRule) GetName() string {
	return r.name
}

func (r *SecondTwelveRule) Eval(rn RouletteNumber) float64 {
	if rn.num > 12 && rn.num < 25 {
		return r.bet * 3
	}

	return 0
}

type ThirdTwelveRule struct {
	name string
	bet  float64
}

func (r *ThirdTwelveRule) GetName() string {
	return r.name
}

func (r *ThirdTwelveRule) Eval(rn RouletteNumber) float64 {
	if rn.num > 24 && rn.num < 37 {
		return r.bet * 3
	}

	return 0
}

type ConcreteNumberRule struct {
	name string
	bet  float64
	num  int
}

func (r *ConcreteNumberRule) GetName() string {
	return r.name
}

func (r *ConcreteNumberRule) Eval(rn RouletteNumber) float64 {
	if rn.num == r.num {
		return r.bet * 36
	}

	return 0
}
