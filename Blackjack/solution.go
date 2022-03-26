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
	case "ten", "jack", "queen", "king":
		return 10
	case "other":
		return 0
	}

	return 0
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	card_one := ParseCard(card1)
	card_two := ParseCard(card2)
	if card_one+card_two == 21 {
		return true
	}
	return false
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack == true {
		if dealerScore < 10 {
			return "W"
		} else {
			return "S"
		}
	}
	return "P"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore >= 17 {
		return "S"
	} else if handScore <= 11 {
		return "H"
	} else if (handScore == 12) || (handScore == 13) || (handScore == 14) || (handScore == 15) || (handScore == 16) {
		if dealerScore <= 6 {
			return "S"
		} else {
			return "H"
		}
	}
	return "0"
}
