package main

import (
	"fmt"
	"github.com/Armingodiz/GoWebDewTraining/10-BlackJackGame/deck"
)

type Dealer struct {
	Cards []deck.Card
}
type Player struct {
	NickName string
	Cards    []deck.Card
}
type User interface {
	userTurn()
	play() int
}

var cards []deck.Card
var holder = 0
var dealer Dealer

func main() {
	cards = deck.NewDeck(deck.MultipleDeck(2), deck.Shuffle)
	fmt.Println(cards)
	dealer = Dealer{[]deck.Card{}}
	dealer.Cards = append(dealer.Cards, cards[holder], cards[holder+1])
	holder += 2
	player := Player{"armin", []deck.Card{}}
	users := []User{}
	users = append(users, &player)
	Start(users)
}

func Start(users []User) {
	for _, user := range users {
		user.userTurn()
	}
	dealer.userTurn()
}
func (dealer *Dealer) userTurn() {
	fmt.Println("DEALER TURN : ")
	fmt.Println(dealer)
}
func (player *Player) userTurn() {
	fmt.Println("PLAYER " + player.NickName + " TURN : ")
	player.Cards = append(player.Cards, cards[holder], cards[holder+1])
	holder += 2
	fmt.Println(player)
	fmt.Println("DEALER CARD IS : " + dealer.Cards[0].Rank.String() + " OF " + dealer.Cards[0].Suit.String() + "s")
	player.play()
}
func (player *Player) play() int { // 0 for stand , 1 for win , -1  for lose
	fmt.Println("1 ) HIT \n2 ) STAND ")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
		player.Cards = append(player.Cards, cards[holder], cards[holder+1])
		holder += 1
		score, _ := scoring(player.Cards)
		if score == 21 {
			return 1
		} else if score > 21 {
			return -1
		} else {
			player.play()
		}
	case 2:
		break
	default:
		fmt.Println("INVALID INPUT !")
		player.play()
	}
	return 0
}
func (dealer *Dealer) play() int {
	fmt.Println("1 ) HIT \n2 ) STAND ")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
	case 2:
		break
	}
	return 0
}
func scoring(cards []deck.Card) (int, string) {
	var score = 0
	var Type = "normal"
	return score, Type
}