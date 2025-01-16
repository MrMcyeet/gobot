package types

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mrmcyeet/gobot/modules/config"
)

type Bot struct {
	Session *discordgo.Session
	Config  *config.Config
}

type Command struct {
	*discordgo.ApplicationCommand

	Execute func(bot *Bot, interaction *discordgo.InteractionCreate) error
}
