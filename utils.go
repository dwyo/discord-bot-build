package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

/*
*

务器: dingwy的服务器 (ID: 1359836604746825748)
分类: 信息 (ID: 1359836605245952121)
文字频道: 欢迎新人 (ID: 1359836605245952122)
文字频道: 笔记-资源 (ID: 1359836605245952123)
分类: 文字频道 (ID: 1359836605245952124)
文字频道: 综合 (ID: 1359836605245952125)
文字频道: 作业-帮助 (ID: 1359836605245952126)
文字频道: 会话-规划 (ID: 1359836605245952127)
文字频道: 题外-话 (ID: 1359836605245952128)
分类: 语音频道 (ID: 1359836605245952129)
语音频道: 休息室 (ID: 1359836605245952130)
语音频道: 自习室 1 (ID: 1359836605652668558)
语音频道: 自习室 2 (ID: 1359836605652668559)
文字频道: 开发 (ID: 1359842430329032814)
*/
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
