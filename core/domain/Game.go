package domain

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"math/rand"
	"sync"
	"time"
)

type Round struct {
	Sets []Set
}
type Game struct {
	Id           string `json:"id"`
	Cards        []Card
	UsedCards    map[string]Card
	CardInGame   map[string]Card
	Players      map[string]Player
	Sets         map[string]Set
	Finish       chan bool
	Pour         chan bool
	Check        chan string
	Pick         chan PickCard
	PickAbleCard int
	DiceNumber   int
	muPick       sync.Mutex
}
type PickCard struct {
	Player string
	Card   string
}

func (g *Game) JoinPlayer(player Player) {
	g.Players[player.Id] = player
}

func (g *Game) Run() {
	for {
		select {
		case <-g.Pour:
			g.DiceNumber = rand.Intn(2)
			g.PickAbleCard = 0

			var numberSet int
			numberSet = 1
			if len(g.Cards) < numberSet {
				numberSet = len(g.Cards)
			}
			for range numberSet {
				index := getRandom(g.Cards)
				if g.Cards[index].Hand == g.DiceNumber {
					g.muPick.Lock()
					g.PickAbleCard++
					g.muPick.Unlock()
					fmt.Println("pick able:", g.PickAbleCard)
				}
				g.CardInGame[g.Cards[index].Id] = g.Cards[index]
				g.Cards = delete_at_card(g.Cards, index)
				fmt.Print("card in game")
				fmt.Println(len(g.Cards) == 0)
				if len(g.Cards) == 0 {
					g.Finish <- true
				}
			}
			rand.Seed(time.Now().UnixNano()) // seed or it will be set to 1
			// generate a random int in the range 0 to 9\
			fmt.Println(g.CardInGame, g.DiceNumber)
		case <-g.Finish:
			var score map[string]int
			score = make(map[string]int)
			for player := range g.Players {
				scorePlayer := len(g.Sets[player].Cards)
				score[player] = scorePlayer
			}
			fmt.Println(score)
		case p := <-g.Pick:
			if card, ok := g.CardInGame[p.Card]; ok {
				if card.Hand == g.DiceNumber {
					g.muPick.Lock()
					g.PickAbleCard--
					g.muPick.Unlock()
					if PlayerSet, ok := g.Sets[p.Player]; ok {
						PlayerSet.Cards = append(PlayerSet.Cards, card)
						g.Sets[p.Player] = PlayerSet

					} else {
						var cards []Card
						cards = append(cards, card)
						g.Sets[p.Player] = Set{Cards: cards}

					}
					delete(g.CardInGame, p.Card)
					if len(g.Cards) == 0 {
						fmt.Println(g.Cards)

					}
				}
			}
		case <-g.Check:
			if g.PickAbleCard == 0 {
				for _, card := range g.CardInGame {
					g.Cards = append(g.Cards, card)
				}
				g.Pour <- true
			}

		}

	}
}

func NewPlayer(name string) *Player {
	return &Player{Id: ulid.Make().String(), Name: name}
}

func NewGame(Cards []Card) *Game {
	return &Game{Id: ulid.Make().String(), Cards: Cards, CardInGame: make(map[string]Card), UsedCards: make(map[string]Card), Players: make(map[string]Player), Sets: make(map[string]Set), Check: make(chan string), Finish: make(chan bool, 10), Pour: make(chan bool, 10), Pick: make(chan PickCard, 10)}
}
func (c *Card) NewCard(color Color, hand int) Card {
	return Card{Id: ulid.Make().String(), Color: color, Hand: hand}
}

func NewCollection() []Card {
	ColorList := []Color{Red}
	var CardList []Card
	for _, color := range ColorList {

		for i := range 2 {
			id := ulid.Make().String()
			CardList = append(CardList, Card{Id: id, Color: color, Hand: i})
		}
	}
	return CardList
}

func getRandom(Card []Card) int {
	rand.Seed(time.Now().UnixNano())    // seed or it will be set to 1
	randomIndex := rand.Intn(len(Card)) // generate a random int in the range 0 to 9
	return randomIndex
}
func delete_at_card(slice []Card, index int) []Card {

	// Append function used to append elements to a slice
	// first parameter as the slice to which the elements
	// are to be added/appended second parameter is the
	// element(s) to be appended into the slice
	// return value as a slice
	return append(slice[:index], slice[index+1:]...)
}
