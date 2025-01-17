package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

		commandList := make([]*discordgo.ApplicationCommand, 0, len(Commands))
		for _, command := range Commands {
			commandList = append(commandList, command.ApplicationCommand)
			fmt.Printf("Attempting to register /%v\n", command.Name)
		}

		if _, err := session.ApplicationCommandBulkOverwrite(session.State.User.ID, "", commandList); err != nil {
			fmt.Println(fmt.Errorf("failed to update commands: %w", err))
		}
	})

	// Setup commands handler
	client.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		//dang, this would've been clever, I think
		//username := map[bool]string{true: interaction.Member.User.Username, false: interaction.User.Username}[interaction.Member.User.Username != ""]

		var username string
		if interaction.Member != nil {
			username = interaction.Member.User.Username
		} else {
			username = interaction.User.Username
		}

		fmt.Printf("@%s executed /%s\n", username, interaction.ApplicationCommandData().Name)

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
	fmt.Printf("@%s is now running. Press CTRL+C to exit.\n", client.State.User.Username)

	Client = &utils.Bot{
		Session: client,
		Config:  config,
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM /*, os.Kill*/)
	<-sc
}
