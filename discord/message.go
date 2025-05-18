package discord

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"os"
)

type Discord struct {
	client webhook.Client
}

func NewDiscord() (*Discord, error) {
	client, err := webhook.NewWithURL(os.Getenv("DISCORD_WEBHOOK"))
	if err != nil {
		return nil, err
	}

	return &Discord{
		client: client,
	}, nil
}

func (d *Discord) PostMessage(msg string, files []*discord.File) error {
	if _, err := d.client.CreateMessage(discord.WebhookMessageCreate{
		Content: msg,
		Files:   files,
	}); err != nil {
		return err
	}
	return nil
}

func (d *Discord) ThatsAll() error {
	return d.PostMessage("That's all for now.", []*discord.File{})
}
