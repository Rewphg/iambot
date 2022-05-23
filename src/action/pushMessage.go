package action

import (
	"os"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func PushMessage(id string, message string) error {
	bot, err := linebot.New(os.Getenv("Channel_Secret"), os.Getenv("Channel_Token"))
	if err != nil {
		return err
	}
	if _, err := bot.PushMessage(id, linebot.NewTextMessage(message)).Do(); err != nil {
		return err
	}

	return nil
}