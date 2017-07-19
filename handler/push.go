package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"time"
)

type PushResource struct {
	db gorm.DB
}

func (pr *PushResource) CreatePushNotification(c *gin.Context) {

}

func (pr *PushResource) GetPushNotificationsBySubscription(c *gin.Context) {
	var notifications []api.Notifications
}

func getPushObjType(obj interface{}) {
	switch obj := obj.(type) {
	case api.Notification:
		return new(api.Notification)
	case api.Subscription:
		return new(api.Subscription)
	default:
		return errors.New("undefined") //put this in api and make a new error that gets caught by sentry?
	}
}
func (pr *PushResource) GetPushObject(c *gin.Context, obj interface{}) {
	id, err := pr.getId(c)
	if err != nil {
		c.JSON(400, api.NewError("problem decoding recv'd id"))
		return
	}
	// api.(obj.(type))

	if err != nil {

	}

	if pr.db.First(&object, id).RecordNotFound() {
		c.JSON(404, gin.H{"error", "not found"}) // later do a new subscription
	} else {
		c.JSON(200, object)
	}
}

func (pr *PushResource) CreatePushSubscription(c *gin.Context) {
	var subscription api.Subscription
	if c.Bind(&subscription) != nil {
		c.JSON(400, errors.New("problem decoding body"))
		return
	}
	subscription.Status = api.SubscriptionStatus
	subscription.created = int32(time.Now().Unix())

	pr.db.Save(&subscription)
	c.JSON(201, subscription)
}

func (pr *TodoResource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}
