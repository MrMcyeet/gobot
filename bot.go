package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	"github.com/mrmcyeet/gobot/modules/config"
	"github.com/mrmcyeet/gobot/modules/testCommand"
	"github.com/mrmcyeet/gobot/modules/utils"
)

var (
	Client   *utils.Bot
	Commands map[string]utils.Command
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	client, err := discordgo.New("Bot " + config.Discord.Token)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
		return
	}

	// Set bot's status :P
	client.AddHandler(func(session *discordgo.Session, _ *discordgo.Ready) {
		session.UpdateStreamingStatus(0, "with Craig 😎", "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	})

	// Register commands
	client.AddHandlerOnce(func(session *discordgo.Session, message *discordgo.Ready) {
		Commands = make(map[string]utils.Command)
		Commands["ping"] = *testCommand.NewPingCommand()

		for _, command := range Commands {
			_, err := session.ApplicationCommandCreate(session.State.User.ID, "1001017041039409233", &discordgo.ApplicationCommand{
				Name:        command.Name,
				Description: command.Description,
			})

			fmt.Printf("Command %s registered\n", command.Name)

			if err != nil {
				fmt.Println(fmt.Errorf("failed to create command %s: %w", command.Name, err))
			}
		}
	})

	// Setup command handlers
	client.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		fmt.Printf("Command get: %s", interaction.ApplicationCommandData().Name)
		command, ok := Commands[interaction.ApplicationCommandData().Name]
		if !ok {
			return
		}

		if err := command.Execute(Client, interaction); err != nil {
			fmt.Println(fmt.Errorf("failed to execute command /%s, invoked by @%s: %w", command.Name, interaction.User.GlobalName, err))
		}
	})

	if err = client.Open(); err != nil {
		log.Fatalf("error opening connection: %v", err)
		return
	}

	defer client.Close()
	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	Client = &utils.Bot{
		Session: client,
		Config:  config,
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt /*syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill*/)
	<-sc
}
