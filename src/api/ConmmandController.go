package api

import (
	"time"

	"github.com/Rewphg/iambot/src/action"
	"github.com/Rewphg/iambot/src/data"
)

func TypeRedirector(c data.EventPost) error {
	for _, j := range c.Event {
		if j.Type == "message" && j.Message.Type == "text" {
			return CommandRedirector(j)
		}
	}
	return nil
}

func CommandRedirector(c data.EventObj) error {
	command := c.Message.Text
	switch command {
	case (".Time"):
		return action.ReplyMessage(c.ReplyToken, time.Now().String())
	case (".Hello"):
		return action.ReplyMessage(c.ReplyToken, "Hello")
	case (".Help"):
		return action.ReplyMessage(c.ReplyToken, ".Time      Get current time\n.Hello      Print Hello\nThis is all command currently.")
	default:
		return action.ReplyMessage(c.ReplyToken, "This command is not in the list. Please use .Help for command list")
	}
}
