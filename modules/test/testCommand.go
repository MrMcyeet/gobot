package testcommand

import (
	"github.com/bwmarrin/discordgo"
	types "github.com/mrmcyeet/gobot/modules/_internal"
)

func NewPingCommand() *types.Command {
	return &types.Command{
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Responds with pong! Used to check if the bot is alive",
		},
		Execute: func(bot *types.Bot, interaction *discordgo.InteractionCreate) error {
			return bot.Session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "Pong! 🏓",
				},
			})
		},
	}
}
