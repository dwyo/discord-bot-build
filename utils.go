package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetAllChannels(s *discordgo.Session) []*discordgo.Channel {
	var channels []*discordgo.Channel
	// 获取机器人加入的所有服务器
	guilds, err := s.UserGuilds(100, "", "", true)
	if err != nil {
		fmt.Println("获取服务器列表失败:", err)
		return nil
	}

	// 遍历每个服务器
	for _, guild := range guilds {
		fmt.Printf("\n服务器: %s (ID: %s)\n", guild.Name, guild.ID)

		// 获取服务器中的所有频道
		channels, err := s.GuildChannels(guild.ID)
		if err != nil {
			fmt.Printf("获取频道列表失败 [%s]: %v\n", guild.Name, err)
			continue
		}

		// 按频道类型分类显示
		for _, channel := range channels {
			switch channel.Type {
			case discordgo.ChannelTypeGuildText:
				fmt.Printf("文字频道: %s (ID: %s)\n", channel.Name, channel.ID)
			case discordgo.ChannelTypeGuildVoice:
				fmt.Printf("语音频道: %s (ID: %s)\n", channel.Name, channel.ID)
			case discordgo.ChannelTypeGuildCategory:
				fmt.Printf("分类: %s (ID: %s)\n", channel.Name, channel.ID)
			}
			channels = append(channels, channel)
		}
	}
	return channels
}
