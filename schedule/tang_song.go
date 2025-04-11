package schedule

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func TangSong(s *discordgo.Session, channelID string) {
	go func() {
		for {
			now := time.Now()
			nextTime := now.Truncate(time.Hour).Add(time.Hour)

			time.Sleep(nextTime.Sub(now))
			// 发送消息到指定频道
			_, err := s.ChannelMessageSend(channelID, "早上好！")
			if err != nil {
				fmt.Println("发送消息时出错:", err)
			}
		}
	}()
}
