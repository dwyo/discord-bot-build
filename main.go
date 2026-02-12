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
	dg.AddHandler(customCmd.MessageCreate)
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := customCmd.CommandHandlers[i.ApplicationCommandData().Name]; ok {
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

	/// 获取所有频道
	// GetAllChannels(dg)
	// 定时任务
	// schedule.TangSong(dg, os.Getenv("SPECIAL_CHANNEL_ID"))

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Printf("Logged in as: %v#%v\n", s.State.User.Username, s.State.User.Discriminator)
}
