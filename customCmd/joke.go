package customCmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

var jokes = []string{
	"为什么程序员总是分不清万圣节和圣诞节？因为 Oct 31 == Dec 25",
	"为什么计算机很怕冷？因为它会冻结（freeze）。",
	"程序员最讨厌的事情是什么？写注释和别人的代码。",
	"一个 SQL 语句走进一家酒吧，走到两张桌子中间，问道：'我可以 join 你们吗？'",
	"为什么程序员喜欢黑暗模式？因为 light attracts bugs（光吸引虫子/bug）。",
	"程序员去买杂货，妻子说'买一袋面包，如果有鸡蛋，买12个'。程序员带着12袋面包回来了。",
	"世界上有10种人：懂二进制的和不懂的。",
	"为什么 Java 程序员戴眼镜？因为他们看不清 C#。",
	"一个前端工程师走进一家店，店员说'我们什么都没有'。前端工程师说'那就给我一个什么都没有的对象吧'。",
	"为什么数据库管理员去不了天堂？因为那里有太多的 DROP。",
}

func Joke(s *discordgo.Session, i *discordgo.InteractionCreate) {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(len(jokes))

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("😄 %s", jokes[result]),
		},
	})
}
