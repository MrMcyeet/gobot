package testCommand

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mrmcyeet/gobot/modules/utils"
)

func NewPingCommand() *utils.Command {

	return utils.NewCommand(
		"ping",
		"Responds with pong! Used to check if the bot is alive",
		func(bot *utils.Bot, interaction *discordgo.InteractionCreate) error {
			return bot.Session.InteractionRespond(interaction.Interaction,
				utils.NewEphemeralResponse("Pong! 🏓"),
			)
		},
	)

}
