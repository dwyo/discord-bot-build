package customCmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Avatar(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	var user *discordgo.User

	if len(options) > 0 {
		user = options[0].UserValue(s)
	} else if i.Member != nil {
		user = i.Member.User
	} else {
		user = i.User
	}

	avatarURL := user.AvatarURL("1024")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("🖼️ %s 的头像:\n%s", user.Username, avatarURL),
		},
	})
}
