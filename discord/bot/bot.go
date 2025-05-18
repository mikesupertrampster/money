package bot

import (
	"github.com/bwmarrin/discordgo"
	"log/slog"
	"regexp"
	"strings"
)

func Start(token string) error {
	goBot, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		return err
	}
	slog.Info("Bot is now connected!")
	return nil
}

func messageHandler(discord *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == discord.State.User.ID {
		return
	}

	re := regexp.MustCompile(`\$([A-Z]{2,})`)
	match := re.FindStringSubmatch(m.Content)

	switch {
	case len(match) > 0:
		getMeGraph(discord, m, match[1])
	case strings.Contains(m.Content, "good bot"):
		if _, err := discord.ChannelMessageSend(m.ChannelID, "You're welcome 😃"); err != nil {
			return
		}
	}
}
