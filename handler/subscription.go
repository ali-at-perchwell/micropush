package handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"micropush/resource/subscription"
)

var (
	updateSubscriptionSql = fmt.Sprintf("UPDATE device_tokens SET token=$1, user_id=$2, push_type=$3, device_id=$4 WHERE device_id=$5")
)

func NewSubscriptionFromJSON(c *gin.Context) (model.Subscription, error) {
	var s subscription.Subscription

	if c.Bind(&person) == nil { // request params
		log.Println(person.Name)
		log.Println(person.Address)
	}
	if c.BindJSON(&s) == nil {
		log.Println(s.Endpoint)
		log.Println(s.Keys)

		var existingToken string
		exists, err := subscription.QueryByAuthToken(s.Keys.Auth)
		if err != nil {
			return nil, err
		}
		if exists {
			// TODO: ADD UPDATE
			// SUBSC TIME ETC

			// or if exists return true, return the subscription object so you dont have to query again?
			return subscription.Update(updateSubscriptionSql, s.Keys.Auth)
		}

		// check to see if subscription exists in the DB

	}

	c.String(200, "Success")
	return subscription // does this need a * or a &

}
