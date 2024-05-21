package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	parsed1, parsed2, dealer := ParseCard(card1), ParseCard((card2)), ParseCard(dealerCard)
	sum := parsed1 + parsed2
	var strat string
	switch {
	case 17 <= sum && sum <= 20, 12 <= sum && sum <= 16 && dealer < 7, sum == 21 && dealer >= 10:
		strat = "S"
	case sum <= 11, 12 <= sum && sum <= 16 && dealer >= 7:
		strat = "H"
	case parsed1 == 11 && parsed2 == 11:
		strat = "P"
	case sum == 21 && dealer <= 10:
		strat = "W"
	}
	return strat
}
