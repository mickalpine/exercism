package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return two
	case "three":
		return three
	case "four":
		return four
	case "five":
		return five
	case "six":
		return six
	case "seven":
		return seven
	case "eight":
		return eight
	case "nine":
		return nine
	case "ten", "jack", "queen", "king":
		return ten
	case "ace":
		return ace
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	dealtScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)
	switch {
	// If you have a pair of aces you must always split them.
	case dealtScore >= 22:
		return split

		// If you have a Blackjack, and dealer does not have 10 or 11, you win.
	case dealtScore == blackjack && (dealerScore != 10 && dealerScore != 11):
		return win

		// If you have a Blackjack, and dealer has 10 or 11 point card then stand.
	case dealtScore == blackjack && dealerScore == 10:
		return stand
	case dealtScore == blackjack && dealerScore == 11:
		return stand

		// If your cards sum up to a value within the range [17, 20] you stand.
	case dealtScore >= 17 && dealtScore <= 20:
		return stand

		// If your cards sum up to a value within the range [12, 16] you stand unless the dealer has a 7 or higher, in which case you should always hit.
	case (dealtScore >= 12 && dealtScore <= 16) && dealerScore < 7:
		return stand
	case (dealtScore >= 12 && dealtScore <= 16) && dealerScore > 7:
		return hit
	default:
		return hit
	}
}

const (
	two       = 2
	three     = 3
	four      = 4
	five      = 5
	six       = 6
	seven     = 7
	eight     = 8
	nine      = 9
	ten       = 10
	jack      = 10
	queen     = 10
	king      = 10
	ace       = 11
	blackjack = 21
)

const (
	hit   = "H"
	split = "P"
	stand = "S"
	win   = "W"
)
