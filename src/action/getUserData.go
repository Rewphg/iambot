package action

import (
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func GetUserData(UID string) (linebot.UserProfileResponse, error) {
	bot, err := linebot.New(os.Getenv("Channel_Secret"), os.Getenv("Channel_Token"))
	if err != nil {
		return linebot.UserProfileResponse{}, err
	}
	res, err := bot.GetProfile(UID).Do()
	if err != nil {
		return linebot.UserProfileResponse{}, err
	}
	return *res, nil
}
