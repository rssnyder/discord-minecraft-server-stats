package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := flag.String("token", "", "discord bot token")
	domain := flag.String("domain", "", "server nickname")
	status := flag.Int("status", 0, "0: playing, 1: listening")
	loop := flag.Int("loop", 60, "seconds between messages")
	flag.Parse()

	dg, err := discordgo.New("Bot " + *token)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	for {

		stats, err := GetMcapiStats(*domain)
		if err != nil {
			log.Printf("%s", err)
			time.Sleep(time.Duration(*loop) * time.Second)
		}

		if stats.Online {

			setActivity(dg, *status, stats.Server.Name)
			time.Sleep(time.Duration(*loop) * time.Second)

			setActivity(dg, *status, fmt.Sprintf("Players: %d/%d", stats.Players.Now, stats.Players.Max))
			time.Sleep(time.Duration(*loop) * time.Second)

			for _, player := range stats.Players.Sample {
				setActivity(dg, *status, player.Name)
				time.Sleep(time.Duration(*loop) * time.Second)
			}
		} else {
			setActivity(dg, *status, "offline")
			time.Sleep(time.Duration(*loop) * time.Second)
		}
	}
}

func setActivity(dg *discordgo.Session, status int, message string) {
	err := dg.UpdateGameStatus(status, message)
	if err != nil {
		log.Printf("Unable to set activity: %s\n", err)
	} else {
		log.Printf("Set activity: %s\n", message)
	}
}
