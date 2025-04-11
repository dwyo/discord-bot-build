package customCmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Calculate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	num1 := options[0].FloatValue()
	operator := options[1].StringValue()
	num2 := options[2].FloatValue()

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "错误：除数不能为0",
				},
			})
			return
		}
		result = num1 / num2
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%.2f %s %.2f = %.2f", num1, operator, num2, result),
		},
	})
}
