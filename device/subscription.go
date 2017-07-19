package device

import (
	"github.com/gin-gonic/gin"
)

type Subscription struct {
	Token   string `json:"token"`
	UserId  string `json:"user_id"`
	Browser string `json:"browser"`
}

// originally called WebToken....
func CheckSubscription(c *gin.Context) error { // still not sure what this method was doing here to begin w/ other than...test..ing
	var s Subscription

	if err := c.BindJSON(&s); err != nil {
		return err // handle the err, this will happen if any of the fields are blank or wrong
	}

	c.String(200, "Success") // or json whichever we want
}

func (sr *SubscriptionResource) Create(sub Subscription) error { // subscription created from json req
	_, err = sr.prepareDb(createSql).Exec(sub.Endpoint, sub.Keys.Auth, time.Now().Unix())

	if err != nil {
		return err
	}
	return nil
}
