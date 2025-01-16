package utils

import "github.com/bwmarrin/discordgo"

// CommandFunc is the function signature for command execution
type CommandFunc func(bot *Bot, interaction *discordgo.InteractionCreate) error

func NewCommand(name string, description string, fn CommandFunc) *Command {
	return &Command{
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        name,
			Description: description,
		},
		Execute: fn,
	}
}
