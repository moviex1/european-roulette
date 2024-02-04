package roulette

import (
	"fmt"
	"time"
)

var menu = []string{
	"",
	"Choose your bets:",
	"-------------------------",
	"1. Even",
	"2. Odd",
	"3. First 18",
	"4. Second 18",
	"5. First 12",
	"6. Second 12",
	"7. Third 12",
	"8. Red",
	"9. Black",
	"10. Concrete number",
	"11. Play",
	"-------------------------",
}

type EuropeanRoulette struct {
	roulette      Roulette
	rouletteTable RouletteTable
}

func BuildEuropeanRoulette() *EuropeanRoulette {
	return &EuropeanRoulette{
		roulette:      BuildRoulette(),
		rouletteTable: BuildRouletteTable(),
	}
}

func (er *EuropeanRoulette) StartGame() {
	balance, name := .0, ""
	choice, bet := 0, .0

	getPlayerInfo(&name, &balance)

	for balance > 0 {
		choice = 0
		for choice != 11 {
			showMenu(balance)
			choice = inputBetween[int]("Your choice: ", 1, 11)

			if choice != 11 {
				bet = inputBetween[float64]("Place your bet: ", 0, balance)
			}

			if rule := handleChoice(choice, bet); rule != nil {
				balance -= bet
				er.rouletteTable.AddRule(rule)
			}

			if choice == 11 && len(er.rouletteTable.rules) == 0 {
				choice = 0
				fmt.Println("You didn't make any bets so you can't play, please place bet, then play")
			}
		}

		spinRouletteAndEvaluateResult(er, &balance)
	}

	fmt.Println("You lost everything :(")
}

func handleChoice(choice int, bet float64) RouletteRule {
	builder := NewRuleBuilder()
	switch choice {
	case 1:
		return builder.SetName(Even).SetBet(bet).Build()
	case 2:
		return builder.SetName(Odd).SetBet(bet).Build()
	case 3:
		return builder.SetName(FirstEighteen).SetBet(bet).Build()
	case 4:
		return builder.SetName(SecondEighteen).SetBet(bet).Build()
	case 5:
		return builder.SetName(FirstTwelve).SetBet(bet).Build()
	case 6:
		return builder.SetName(SecondTwelve).SetBet(bet).Build()
	case 7:
		return builder.SetName(ThirdTwelve).SetBet(bet).Build()
	case 8:
		return builder.SetName(RedColor).SetBet(bet).Build()
	case 9:
		return builder.SetName(BlackColor).SetBet(bet).Build()
	case 10:
		num := inputBetween[int]("Input number(from 0 to 36) that you want to place bet: ", 0, 36)
		return builder.SetName(ConcreteNumber).SetNum(num).SetBet(bet).Build()
	}
	return nil
}

func getPlayerInfo(name *string, balance *float64) {
	fmt.Print("Input your name: ")
	fmt.Scan(name)
	fmt.Print("Input your balance: ")
	fmt.Scan(balance)
}

func showMenu(balance float64) {
	for _, str := range menu {
		fmt.Println(str)
	}

	fmt.Printf("Your balance: %g \n\n", balance)
}

func spinRouletteAndEvaluateResult(er *EuropeanRoulette, balance *float64) {
	defer er.rouletteTable.ClearTable()
	rouletterNumber := er.roulette.Play()

	fmt.Println("Andddd... roulette starts to spinnig...")
	time.Sleep(2 * time.Second)
	fmt.Printf("And we have %d %s !\n", rouletterNumber.num, rouletterNumber.color)
	time.Sleep(1 * time.Second)

	win := er.rouletteTable.CalculateWin(rouletterNumber)
	er.rouletteTable.DisplayBetsResult(rouletterNumber)

	if win == 0 {
		fmt.Println("You lose :(")
	} else {
		fmt.Printf("You won %g!\n", win)
	}

	*balance += win
}

func inputBetween[V int | float64](text string, from, to V) V {
	var value V
	fmt.Print(text)
	fmt.Scan(&value)
	for value < from || value > to {
		fmt.Print(text)
		fmt.Scan(&value)
	}

	return value
}
