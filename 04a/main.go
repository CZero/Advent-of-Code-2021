package main

import (
	"fmt"
	"strconv"
	"strings"
)

type coords struct {
	r int // Row
	c int // Col
}
type bingocard map[coords]string // bingocard contains a bingocard with lines like 1,1=12 1,2=4, 1,3=45 etc denoting the numbers.
type cards []bingocard           // cards is a collection of bingocards

var (
	NBingoCards    int             // nBingocards keeps track of the amount of cards in play.
	WinningNumbers []string        // Duh
	GameRound      int             // The game progress, the round we are in
	PulledNumbers  map[string]bool // Keeps track of the numbers that were pulled
	Cards          cards           // The cards
	RowsPerCard    int             // The numbers of row per card
	ColsPerCard    int             // The number of cols per card
)

func main() {
	// lines, _ := readLines("example.txt")
	// lines, _ := readLines("example2.txt")
	lines, _ := readLines("input.txt")
	CreateCards(lines)
	fmt.Println(PlayGame())
}

// createCards walks through the lines, gets the winning numbers and creates the cards.
func CreateCards(input []string) {
	// First we get the winning numbers.
	WinningNumbers = strings.Split(input[0], ",")

	// Then we get the cards. An empty line is end of a card
	var row int // Row on active card
	for i := 1; i < len(input); i++ {
		if input[i] == "" { // An empty lines means another card is coming, the previous is done!
			NBingoCards++
			card := make(bingocard)
			Cards = append(Cards, card)
			RowsPerCard = row
			row = 0
		} else { // A line with content, we found another row of numbers for this card!
			numbers := strings.Fields(input[i])
			for j := 0; j < len(numbers); j++ {
				coord := coords{row, j}
				Cards[NBingoCards-1][coord] = numbers[j]
			}
			row++
			ColsPerCard = len(numbers)
		}
	}
}

// PlayGame 🎼 Play the game
func PlayGame() string {
	var winners string
	PulledNumbers = make(map[string]bool)
	for ; GameRound < len(WinningNumbers); GameRound++ {
		PulledNumbers[WinningNumbers[GameRound]] = true
		for i, card := range Cards {
			for r := 0; r < RowsPerCard; r++ {
				for c := 0; c < ColsPerCard; c++ {
					if card[coords{r, c}] == WinningNumbers[GameRound] {
						if GameRound > 3 { // Before gameround 4 there aren't 5 winning numbers, so why check?
							if CheckForWinner(coords{r, c}, card, PulledNumbers) { // Don't return just yet, what if multiple cards won??
								winners += fmt.Sprintf("Card %d won with winning number %s in round %d.\nThe score: %d\n", i, WinningNumbers[GameRound], GameRound, CalculateScore(card, PulledNumbers, WinningNumbers[GameRound]))
							}
						}
					}
				}
			}
		}
		if winners != "" {
			DisplayCards(PulledNumbers) // Show the cards when someone won
			return winners              // Returns all the winners
		}
	}
	DisplayCards(PulledNumbers) // Show the cards when all lost
	return "No winner, no chicken dinner"
}

// CheckForWinner checks if there's a winning card based on the pulled number (coords!)
func CheckForWinner(Coords coords, card bingocard, pulled map[string]bool) bool {
	// check Row
	for col, winners := 0, 0; col < ColsPerCard; col++ {
		if pulled[card[coords{Coords.r, col}]] {
			winners++
			if winners == ColsPerCard {
				return true
			}
		}
	}
	for row, winners := 0, 0; row < RowsPerCard; row++ {
		if pulled[card[coords{row, Coords.c}]] {
			winners++
			if winners == RowsPerCard {
				return true
			}
		}
	}
	// check Col
	return false
}

// CalculateScore takes the card and the pulled numbers and winning number. It sums the unpulled numbers and multiplies by winnning number.
func CalculateScore(card bingocard, pulled map[string]bool, winningnumber string) (score int) {
	winner, _ := strconv.Atoi(winningnumber)
	for r := 0; r < RowsPerCard; r++ {
		for c := 0; c < ColsPerCard; c++ {
			if !pulled[card[coords{r, c}]] {
				number, _ := strconv.Atoi(card[coords{r, c}])
				score += number
			}
		}
	}
	return score * winner
}

// DisplayCards presents the cards in a human readable format, together with the pulled numbers, for diagnostics.
func DisplayCards(pulled map[string]bool) {
	for n, card := range Cards {
		fmt.Println("Cardnumber: ", n)
		for r := 0; r < RowsPerCard; r++ {
			var row string
			for c := 0; c < ColsPerCard; c++ {
				if len(card[coords{r, c}]) == 2 {
					if pulled[card[coords{r, c}]] {
						row += "*" + card[coords{r, c}] + "* "
					} else {
						row += " " + card[coords{r, c}] + "  "
					}
				}
				if len(card[coords{r, c}]) == 1 {
					if pulled[card[coords{r, c}]] {
						row += "*" + card[coords{r, c}] + " * "
					} else {
						row += " " + card[coords{r, c}] + "   "
					}
				}
			}
			fmt.Println(row)
		}
		fmt.Println()
	}
}
