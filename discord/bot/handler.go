package bot

import (
	"awesomeProject/intelligence"
	"awesomeProject/types"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log/slog"
	"os"
)

func getGraph(symbol string) (*os.File, error) {
	f := intelligence.Finviz{}
	if err := intelligence.New(&f); err != nil {
		return nil, err
	}
	path, err := f.GetGraph(symbol)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getMetrics(symbol string) (types.Metrics, error) {
	var m types.Metrics

	f := intelligence.Finviz{}
	if err := intelligence.New(&f); err != nil {
		return m, err
	}
	m, err := f.GetMetrics(symbol)
	if err != nil {
		return m, err
	}
	return m, nil

}

func Oops(discord *discordgo.Session, channelID string) error {
	if _, err := discord.ChannelMessageSend(channelID, "Oops, something went wrong... :crying_cat_face:"); err != nil {
		return err
	}
	return nil
}

func getMeGraph(discord *discordgo.Session, m *discordgo.MessageCreate, symbol string) {
	msg := fmt.Sprintf("I see you have mentioned **%s**, I'm getting the graph for you...", symbol)
	if _, err := discord.ChannelMessageSend(m.ChannelID, msg); err != nil {
		slog.Error(err.Error())
		return
	}

	graph, err := getGraph(symbol)
	if err != nil {
		slog.Error(err.Error())
		err = Oops(discord, m.ChannelID)
		if err != nil {
			slog.Error(err.Error())
		}
		return
	}

	metrics, err := getMetrics(symbol)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if _, err := discord.ChannelMessageSendComplex(
		m.ChannelID,
		&discordgo.MessageSend{
			Content: "",
			Files:   []*discordgo.File{{Name: "chart.png", Reader: graph}},
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeRich,
				Title:       "Current Info",
				Description: fmt.Sprintf("Ticker: [%[1]s](<https://finviz.com/quote.ashx?t=%[1]s&ty=c&p=d&b=1>)", symbol),
				Color:       0x1c4399,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Market Cap",
						Value:  metrics.MarketCap,
						Inline: true,
					},
					{
						Name:   "Price",
						Value:  "$" + metrics.Price,
						Inline: true,
					},
					{
						Name:   "Change",
						Value:  metrics.Change,
						Inline: true,
					},
					{
						Name:   "Volume",
						Value:  metrics.Volume,
						Inline: true,
					},
					{
						Name:   "52W Range",
						Value:  metrics.FiftyTwoWeekRange,
						Inline: true,
					},
					{
						Name:   "52W High",
						Value:  metrics.FiftyTwoWeekHigh,
						Inline: true,
					},
					{
						Name:   "52W Low",
						Value:  metrics.FiftyTwoWeekLow,
						Inline: true,
					},
					{
						Name:   "RSI (14)",
						Value:  metrics.RSI14,
						Inline: true,
					},
				},
			}},
		},
	); err != nil {
		slog.Error(err.Error())
	}
}
