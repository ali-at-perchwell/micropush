package push

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/googlechrome/push-encryption-go/webpush"
	sub "micropush/resource/web/subscription"
)

// type Pusher interface {
// 	Push() error
// }

// type WebPush struct {
// 	Sender   sub.Subscription
// 	Receiver sub.Subscription
// 	Message  string
// 	VapidKey string
// }

// type ApnsPushEvent struct {
// }

func Push(sender sub.Subscription, msg string, vKey string) error {
	// endpt would be the receiver
	sub := &webpush.Subscription{sender.Endpoint, sender.Keys.P256DH, sender.Keys.Auth}
	res, err := webpush.Send(nil, sub, msg, vKey) // TODO: check the msg for illegal stuff and vKey
}

func Create(c *gin.Context) { // get it from the req params now

}
