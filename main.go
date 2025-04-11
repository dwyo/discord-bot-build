package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"dwyo/discord-bot/customCmd"
)

var (
	commands = []*discordgo.ApplicationCommand{
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
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping":      customCmd.Ping,
		"hello":     customCmd.Hello,
		"goodbye":   customCmd.Goodbye,
		"calculate": customCmd.Calculate,
	}
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fmt.Println("Error: BOT_TOKEN not found in environment. Please set your bot token.")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	// 设置必要的 Intents
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	// 注册斜杠命令
	appID := os.Getenv("BOT_APP_ID")
	if appID == "" {
		fmt.Println("Error: APPLICATION_ID not found in environment")
		return
	}

	// fmt.Println("正在注册斜杠命令...")
	// for _, v := range commands {
	// 	_, err := dg.ApplicationCommandCreate(appID, "", v)
	// 	if err != nil {
	// 		fmt.Printf("注册命令 %v 时出错: %v\n", v.Name, err)
	// 		return
	// 	}
	// }

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Printf("Logged in as: %v#%v\n", s.State.User.Username, s.State.User.Discriminator)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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
		for _, v := range commands {
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
