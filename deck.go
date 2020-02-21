package main

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
)

type Deck struct {
	cards [78]Card
	layouts [][]int
}

type Card struct {
	Name    string
	Meaning []string
}

func (d *Deck) SelectRandomCard() Card {
	return d.cards[rand.Intn(len(d.cards))]
}

func CreateMessagesInLayout(cards []Card, layout []int) []discordgo.MessageSend {
	var messages []discordgo.MessageSend

	for _, numCardsInMessage := range layout {
		var cardsInMessage []Card
		for i := 0; i < numCardsInMessage; i++ {
			var card Card
			card, cards = PopCard(cards)
			cardsInMessage = append(cardsInMessage, card)
		}
		messages = append(messages, CreateMessageWithCards(cardsInMessage))
	}

	return messages
}

func CreateMessageWithCards(cards []Card) discordgo.MessageSend {
	var content []string
	for _, card := range cards {
		content = append(content, card.Name)
	}
	return discordgo.MessageSend{Content: strings.Join(content, " ")}
}

func PopCard(a []Card) (Card, []Card) {
	return a[0], a[1:]
}

func NewDeck() Deck {
	return Deck{
		layouts: [][]int{
			[]int{1},
			[]int{2},
			[]int{3},
			[]int{4},
			[]int{1,3,1},
			[]int{1,2,3},
			[]int{3,1,3},
			[]int{1,2,2,2,1},
			[]int{3,3,3},
			[]int{1,4,3,4},
		},
		cards: [78]Card{
			Card{Name: "The Fool", Meaning: []string{}},
			Card{Name: "The Magician", Meaning: []string{}},
			Card{Name: "The High Priestess", Meaning: []string{}},
			Card{Name: "The Empress", Meaning: []string{}},
			Card{Name: "The Emperor", Meaning: []string{}},
			Card{Name: "The Hierophant", Meaning: []string{}},
			Card{Name: "The Lovers", Meaning: []string{}},
			Card{Name: "The Chariot", Meaning: []string{}},
			Card{Name: "Strength", Meaning: []string{}},
			Card{Name: "The Hermit", Meaning: []string{}},
			Card{Name: "Wheel of Fortune", Meaning: []string{}},
			Card{Name: "Justice", Meaning: []string{}},
			Card{Name: "The Hanged Man", Meaning: []string{}},
			Card{Name: "Death", Meaning: []string{}},
			Card{Name: "Temperance", Meaning: []string{}},
			Card{Name: "The Devil", Meaning: []string{}},
			Card{Name: "The Tower", Meaning: []string{}},
			Card{Name: "The Star", Meaning: []string{}},
			Card{Name: "The Moon", Meaning: []string{}},
			Card{Name: "The Sun", Meaning: []string{}},
			Card{Name: "Judgement", Meaning: []string{}},
			Card{Name: "The World", Meaning: []string{}},
			Card{Name: "Ace of Wands", Meaning: []string{}},
			Card{Name: "Two of Wands", Meaning: []string{}},
			Card{Name: "Three of Wands", Meaning: []string{}},
			Card{Name: "Four of Wands", Meaning: []string{}},
			Card{Name: "Five of Wands", Meaning: []string{}},
			Card{Name: "Six of Wands", Meaning: []string{}},
			Card{Name: "Seven of Wands", Meaning: []string{}},
			Card{Name: "Eight of Wands", Meaning: []string{}},
			Card{Name: "Nine of Wands", Meaning: []string{}},
			Card{Name: "Ten of Wands", Meaning: []string{}},
			Card{Name: "Page of Wands", Meaning: []string{}},
			Card{Name: "Knight of Wands", Meaning: []string{}},
			Card{Name: "Queen of Wands", Meaning: []string{}},
			Card{Name: "King of Wands", Meaning: []string{}},
			Card{Name: "Ace of Cups", Meaning: []string{}},
			Card{Name: "Two of Cups", Meaning: []string{}},
			Card{Name: "Three of Cups", Meaning: []string{}},
			Card{Name: "Four of Cups", Meaning: []string{}},
			Card{Name: "Five of Cups", Meaning: []string{}},
			Card{Name: "Six of Cups", Meaning: []string{}},
			Card{Name: "Seven of Cups", Meaning: []string{}},
			Card{Name: "Eight of Cups", Meaning: []string{}},
			Card{Name: "Nine of Cups", Meaning: []string{}},
			Card{Name: "Ten of Cups", Meaning: []string{}},
			Card{Name: "Page of Cups", Meaning: []string{}},
			Card{Name: "Knight of Cups", Meaning: []string{}},
			Card{Name: "Queen of Cups", Meaning: []string{}},
			Card{Name: "King of Cups", Meaning: []string{}},
			Card{Name: "Ace of Swords", Meaning: []string{}},
			Card{Name: "Two of Swords", Meaning: []string{}},
			Card{Name: "Three of Swords", Meaning: []string{}},
			Card{Name: "Four of Swords", Meaning: []string{}},
			Card{Name: "Five of Swords", Meaning: []string{}},
			Card{Name: "Six of Swords", Meaning: []string{}},
			Card{Name: "Seven of Swords", Meaning: []string{}},
			Card{Name: "Eight of Swords", Meaning: []string{}},
			Card{Name: "Nine of Swords", Meaning: []string{}},
			Card{Name: "Ten of Swords", Meaning: []string{}},
			Card{Name: "Page of Swords", Meaning: []string{}},
			Card{Name: "Knight of Swords", Meaning: []string{}},
			Card{Name: "Queen of Swords", Meaning: []string{}},
			Card{Name: "King of Swords", Meaning: []string{}},
			Card{Name: "Ace of Coins", Meaning: []string{}},
			Card{Name: "Two of Coins", Meaning: []string{}},
			Card{Name: "Three of Coins", Meaning: []string{}},
			Card{Name: "Four of Coins", Meaning: []string{}},
			Card{Name: "Five of Coins", Meaning: []string{}},
			Card{Name: "Six of Coins", Meaning: []string{}},
			Card{Name: "Seven of Coins", Meaning: []string{}},
			Card{Name: "Eight of Coins", Meaning: []string{}},
			Card{Name: "Nine of Coins", Meaning: []string{}},
			Card{Name: "Ten of Coins", Meaning: []string{}},
			Card{Name: "Page of Coins", Meaning: []string{}},
			Card{Name: "Knight of Coins", Meaning: []string{}},
			Card{Name: "Queen of Coins", Meaning: []string{}},
			Card{Name: "King of Coins", Meaning: []string{}},
		},
	}
}
