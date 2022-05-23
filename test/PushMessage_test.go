package action

import (
	"testing"

	"github.com/Rewphg/iambot/src/action"
)

func TestSendingImageMessage(t *testing.T) {
	if err := action.PushImageMessage("U8a48f9b33da301fe5fbeb8f351b0bd25", "/mnt/Synology/BlueIris/Picture/TestPicture.png", "/mnt/Synology/BlueIris/Picture/TestPicture.png"); err != nil {
		t.Fatalf(`Error Sending Picture : < %s >`, err)
	}
}
