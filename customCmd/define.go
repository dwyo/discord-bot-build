package customCmd

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "回复 Pong!",
		},
		{
			Name:        "hello",
			Description: "打个招呼",
		},
		{
			Name:        "goodbye",
			Description: "说再见",
		},
		{
			Name:        "calculate",
			Description: "进行四则运算",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionNumber,
					Name:        "num1",
					Description: "第一个数字",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "operator",
					Description: "运算符 (+, -, *, /)",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{Name: "加法", Value: "+"},
						{Name: "减法", Value: "-"},
						{Name: "乘法", Value: "*"},
						{Name: "除法", Value: "/"},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionNumber,
					Name:        "num2",
					Description: "第二个数字",
					Required:    true,
				},
			},
		},
		{
			Name:        "roll",
			Description: "掷一个六面骰子",
		},
		{
			Name:        "avatar",
			Description: "获取用户头像",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "要查看头像的用户（可选，不填则显示自己的）",
					Required:    false,
				},
			},
		},
		{
			Name:        "random",
			Description: "生成指定范围内的随机数",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "min",
					Description: "最小值",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "max",
					Description: "最大值",
					Required:    true,
				},
			},
		},
		{
			Name:        "joke",
			Description: "获取一个随机程序员笑话",
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping":      Ping,
		"hello":     Hello,
		"goodbye":   Goodbye,
		"calculate": Calculate,
		"roll":      Roll,
		"avatar":    Avatar,
		"random":    Random,
		"joke":      Joke,
	}
)

// MessageCreate 处理消息创建事件，用于同步命令
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!synccommands" {
		appID := os.Getenv("BOT_APP_ID")
		if appID == "" {
			s.ChannelMessageSend(m.ChannelID, "错误：未设置 BOT_APP_ID 环境变量")
			return
		}

		// 删除所有现有命令
		registerCommands, err := s.ApplicationCommands(appID, "")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "获取现有命令时出错："+err.Error())
			return
		}
		fmt.Println("delete commands:")
		for _, cmd := range registerCommands {
			fmt.Println(cmd.Name)
			err := s.ApplicationCommandDelete(appID, "", cmd.ID)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("删除命令 %s 时出错：%v", cmd.Name, err))
				return
			}
		}
		fmt.Println("delete commands end")

		fmt.Println("register commands:")
		// 重新注册所有命令
		for _, v := range Commands {
			fmt.Println(v.Name)
			_, err := s.ApplicationCommandCreate(appID, "", v)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("注册命令 %s 时出错：%v", v.Name, err))
				return
			}
		}
		fmt.Println("register commands end")

		s.ChannelMessageSend(m.ChannelID, "命令同步完成！")
	}
}
