package customCmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
)

func Hello(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// spew.Dump(s)
	spew.Dump(i)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello!",
		},
	})
}
