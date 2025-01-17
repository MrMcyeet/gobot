package utils

import "github.com/bwmarrin/discordgo"

// CommandFunc is the function signature for command execution
type CommandFunc func(bot *Bot, interaction *discordgo.InteractionCreate) error

func NewSimpleCommand(name string, description string, fn CommandFunc) *Command {
	return &Command{
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        name,
			Description: description,
		},
		Execute: fn,
	}
}

func NewUserCommand(name string, description string, fn CommandFunc) *Command {
	return &Command{
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:             name,
			Description:      description,
			IntegrationTypes: &[]discordgo.ApplicationIntegrationType{discordgo.ApplicationIntegrationUserInstall},
			Contexts:         &[]discordgo.InteractionContextType{discordgo.InteractionContextGuild, discordgo.InteractionContextBotDM, discordgo.InteractionContextPrivateChannel},
		},
		Execute: fn,
	}
}
