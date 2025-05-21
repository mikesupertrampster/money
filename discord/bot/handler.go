package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log/slog"
	intelligence "money/intelligence/scrape"
	"os"
)

func getMe(discord *discordgo.Session, m *discordgo.MessageCreate, symbol string) {
	msg := fmt.Sprintf("Getting **%s**...", symbol)
	if _, err := discord.ChannelMessageSend(m.ChannelID, msg); err != nil {
		slog.Error(err.Error())
		return
	}

	finviz, err := intelligence.NewFinviz()
	if err != nil {
		slog.Error(err.Error())
	}
	obj, err := finviz.GetMetrics(symbol)
	if err != nil {
		return
	}
	img, err := os.Open(obj.Image)

	if _, err = discord.ChannelMessageSendComplex(
		m.ChannelID,
		&discordgo.MessageSend{
			Content: "",
			Files:   []*discordgo.File{{Name: "chart.png", Reader: img}},
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeRich,
				Title:       "Current Info",
				Description: fmt.Sprintf("Ticker: [%[1]s](<https://finviz.com/quote.ashx?t=%[1]s&ty=c&p=d&b=1>)", symbol),
				Color:       0x1c4399,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Market Cap",
						Value:  obj.MarketCap,
						Inline: true,
					},
					{
						Name:   "Price",
						Value:  "$" + obj.Price,
						Inline: true,
					},
					{
						Name:   "Change",
						Value:  obj.Change,
						Inline: true,
					},
					{
						Name:   "Volume",
						Value:  obj.Volume,
						Inline: true,
					},
					{
						Name:   "52W Range",
						Value:  obj.FiftyTwoWeekRange,
						Inline: true,
					},
					{
						Name:   "52W High",
						Value:  obj.FiftyTwoWeekHigh,
						Inline: true,
					},
					{
						Name:   "52W Low",
						Value:  obj.FiftyTwoWeekLow,
						Inline: true,
					},
					{
						Name:   "RSI (14)",
						Value:  obj.RSI14,
						Inline: true,
					},
				},
			}},
		},
	); err != nil {
		slog.Error(err.Error())
	}
}
