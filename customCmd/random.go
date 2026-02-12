package customCmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Random(s *discordgo.Session, i *discordgo.InteractionCreate) {
	rand.Seed(time.Now().UnixNano())
	options := i.ApplicationCommandData().Options
	min := int(options[0].IntValue())
	max := int(options[1].IntValue())

	if min > max {
		min, max = max, min
	}

	result := rand.Intn(max-min+1) + min

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("🔢 随机数 (%d-%d): %d", min, max, result),
		},
	})
}
