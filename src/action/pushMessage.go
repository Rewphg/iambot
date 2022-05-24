package action

import (
	"fmt"
	"os"

	ix "github.com/imgix/imgix-go/v2"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func PushTextMessage(id string, message string) error {
	bot, err := linebot.New(os.Getenv("Channel_Secret"), os.Getenv("Channel_Token"))
	if err != nil {
		return err
	}
	if _, err := bot.PushMessage(id, linebot.NewTextMessage(message)).Do(); err != nil {
		return err
	}

	return nil
}

func PushImageMessage(id string, original_path string, preview_path string) error {
	ub := ix.NewURLBuilder("www.chokrporn.one", ix.WithLibParam(false))
	ixURL := ub.CreateURL(original_path)
	fmt.Println(ixURL)

	bot, err := linebot.New("c69397876c388203269ca0f71cb171fb", "IQQDzJx/sfhbSulfgCCFlzCTU860hzWkeF9YAeRLnPRzjLdZwt4znFyBvbANej0rzsPjRPoRgKou+fk2eiA6G5LzB4wL1TF6UDbYQsJwzRTxp5Ie7DVWjrM5EJfIP4O3H9yYiM/U3PN6wNlTJnZtkgdB04t89/1O/w1cDnyilFU=")
	if err != nil {
		return err
	}
	if _, err := bot.PushMessage(id, linebot.NewImageMessage(ixURL, ixURL)).Do(); err != nil {
		return err
	}

	return nil
}
