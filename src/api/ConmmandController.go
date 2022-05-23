package api

import (
	"time"
	"log"

	"github.com/Rewphg/iambot/src/action"
	"github.com/Rewphg/iambot/src/data"
)

func TypeRedirector(c data.EventPost) error {
	for _, j := range c.Event {
		switch j.Source.Type {
		case ("user"):
			log.Println("Message Sent From User")
			return	UserSwitches(j)
		case ("group"):
			log.Println("Message Sent From Group")
			return	UserSwitches(j)
		case ("room"):
			log.Println("Message Sent From Room")
			return	action.ReplyMessage(j.ReplyToken, "Room command is not aviable")
		default:
			
		}
	}
	return nil
}

func UserSwitches(c data.EventObj) error {

	if c.Type != "message" && c.Message.Type != "text" {
		return action.ReplyMessage(c.ReplyToken, "This operation is invalid")
	}

	command := c.Message.Text
	switch command {
		case (".Time"):
			return action.ReplyMessage(c.ReplyToken, time.Now().String())
		case (".Hello"):
			return action.ReplyMessage(c.ReplyToken, "Hello")
		case (".Help"):
			return action.ReplyMessage(c.ReplyToken, ".Time      Get current time\n.Hello      Print Hello\nThis is all command currently.")
		default:
			return nil
	}
}
