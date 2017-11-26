package main

import (
	"github.com/turnage/graw/reddit"
	"fmt"
	"time"
	"github.com/turnage/graw"
    "github.com/bwmarrin/discordgo"
)
type announcer struct{}

func (a *announcer) Post(post *reddit.Post) error {
	fmt.Printf(`%s posted "%s"\n`, post.Author, post.Title)
	return nil
}

func main() {

	discord, errDiscord := discordgo.New("authentication token")
	if errDiscord != nil {
		fmt.Printf("Error while initializing discord API ")
	}
	fmt.Printf("Launching Script !")
	// Get an api handle to reddit for a logged out (script) program,
	// which forwards this user agent on all requests and issues a request at
	// most every 5 seconds.
	apiHandle, _ := reddit.NewScript("VirtualFox", 5 * time.Second)

	// Create a configuration specifying what event sources on Reddit graw
	// should connect to the bot.
	cfg := graw.Config{Subreddits: []string{"lyon","gaming"}}

	// launch a graw scan in a goroutine using the bot, handle, and config. The
	// returned "stop" and "wait" are functions. "stop" will stop the graw run
	// at any time, and "wait" will block until it finishes.
	_, wait, _ := graw.Scan(&announcer{}, apiHandle, cfg)

	// This time, let's block so the bot will announce (ideally) forever.
	if err := wait(); err != nil {
		fmt.Printf("graw run encountered an error: %v\n", err)
	}
}
