package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var authToken string
var callsign = ".gotarot"
var maxNumCards = 10
var deck Deck

func init() {
	flag.StringVar(&authToken, "token", "", "The authentication token for your Discord server")
	flag.Parse()
	deck = NewDeck()
}

func main() {

	// Make sure the Auth Token was passed
	if authToken == "" {
		panic("Missing --token")
		//fmt.Println("Missing --token <token>")
		//fmt.Println("To get a new token ")
		//// Create a new Discord session using the provided login information.
		//dg, err := discordgo.New(Email, Password)
		//if err != nil {
		//	fmt.Println("error creating Discord session,", err)
		//	return
		//}
		//
		//// Print out your token.
		//fmt.Printf("Your Authentication Token is:\n\n%s\n", dg.Token)
	}

	// Sign in to the discord server
	dg, err := discordgo.New("Bot " + authToken)
	if err != nil {
		panic(err)
	}
	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("GoTarot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateStatus(0, callsign)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// check if the message is the callsign
	if strings.HasPrefix(m.Content, callsign) {

		// Determine number of cards
		content_slice := strings.Split(m.Content, " ")
		numCards, err := strconv.Atoi(content_slice[len(content_slice) - 1])
		if err != nil {
			// Last element in slice is not an integer
			_, err := s.ChannelMessageSend(m.ChannelID, help())
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}
		if numCards > maxNumCards {
			// Too many cards to draw
			_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprint("I can only draw up to %i cards!", maxNumCards))
			if err != nil {
				fmt.Println(err.Error())
			}

		}

		// Select cards
		var drawnCards []Card
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < numCards; i++ {
			drawnCards = append(drawnCards, deck.SelectRandomCard())
		}

		// Get layout
		layout := deck.layouts[numCards - 1]
		messages := CreateMessagesInLayout(drawnCards, layout)

		// Send cards to channel
		for _, message := range messages {
			_, err = s.ChannelMessageSendComplex(m.ChannelID, &message)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func help() string {
	return "I read tarot, not minds!  To draw cards use the format !gotarot <num of cards>"
}
