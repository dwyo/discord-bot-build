package customCmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Hello(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// spew.Dump(s)
	/**
	i.Member信息:
	(*discordgo.Member)(0x140000cbb80)({
			GuildID: (string) "",
			JoinedAt: (time.Time) 2025-04-10 10:25:11.141 +0000 +0000,
			Nick: (string) "",
			Deaf: (bool) false,
			Mute: (bool) false,
			Avatar: (string) "",
			User: (*discordgo.User)(0x140000c95f0)(dingwy),
			Roles: ([]string) {
		},
		PremiumSince: (*time.Time)(<nil>),
		Flags: (discordgo.MemberFlags) 0,
		Pending: (bool) false,
		Permissions: (int64) 2251799813685247,
		CommunicationDisabledUntil: (*time.Time)(<nil>)
	})
	*/

	// spew.Dump(i.Member.User.Username)

	var username string
	if i.Member != nil {
		username = i.Member.User.Username
	} else {
		username = i.User.Username
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Hello, %s!", username),
		},
	})
}
