package model

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	//"microservice/service" for err reporting??
)

type Subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     string `json:"keys"`
}

type Keys struct {
	P256DH string `json:"p256dh"`
	Auth   string `json:"auth"`
}

var ( // TODO: CHANGE SQL
	createSql = fmt.Sprintf("INSERT INTO tokens_web (token, user_id, created_at) VALUES ($1, $2, $3)")
	updateSql = fmt.Sprintf("UPDATE device_tokens SET token=$1, user_id=$2, push_type=$3, device_id=$4 WHERE device_id=$5")
	deleteSql = fmt.Sprintf("UPDATE device_tokens SET token=$1, user_id=$2, push_type=$3, device_id=$4 WHERE device_id=$5")
)

type SubscriptionResource struct {
	db sql.DB // debating this
}

func (sr *SubscriptionResource) prepareDb(rawText string) (sqlTx sql.Tx) {
	sqlTx, err := sr.db.Prepare(rawText)
	if err != nil {
		raven.CaptureError(err, nil)
		return nil, err
	}
	defer sqlTx.Close() // yes?

	return sqlTx, nil
}

// NOTE: NOT SUPER DRY...WILL REVISIT

// BEGIN???? OR JUST PREP
func (sr *SubscriptionResource) CreateSubscription(c *gin.Context) { // subscription created from json req
	var s Subscription
	if err := c.BindJSON(&s); err != nil {
		c.JSON(400, errors.New("Err: Problem decoding id sent")) // or send a specific one...nidids
		return
	}
	_, err = sr.prepareDb(createSql).Exec(s.Endpoint, s.Keys.Auth, time.Now().Unix())

	if err != nil {
		c.JSON(400, errors.New("Err: Problem querying the db")) // or send a specific one...nidids
		return
	}
	c.JSON(200, s)
}

func (sr *SubscriptionResource) GetSubscription(c *gin.Context) Subscription {
	var s Subscription
	if err := c.BindJSON(&s); err != nil {
		c.JSON(400, errors.New("Err: Problem decoding id sent")) // or send a specific one...nidids
		return nil
	} // ret err?????
	return s
}

func (sr *SubscriptionResource) DeleteSubscription(c *gin.Context) {
	s := GetSubscription(c) // errrrr

	_, err = sr.prepareDb(deleteSql).Exec(s.Endpoint, s.Keys.Auth, time.Now().Unix())
	if err != nil {
		c.JSON(400, errors.New("Err: Problem deleting subscription"))
		return
	}
	c.JSON(200, "'msg': 'subscription deleted'")
}

func (sr *SubscriptionResource) UpdateSubscription(c *gin.Context) {
	s := GetSubscription(c) // errrrr

	_, err = sr.prepareDb(updateSql).Exec(sub.Endpoint, sub.Keys.Auth, time.Now().Unix())
	if err != nil {
		c.JSON(400, errors.New("Err: Problem updating subscription")) // check the codes
		return
	}
	c.JSON(200, "updated")
}

func (sr *SubscriptionResource) getSubscriptionByAuthToken(c *gin.Context) (Subscription, error) { //*?????
	authStr := c.Params.ByName("auth") // why would they have teh id
	var s subscription
	err := sr.QueryRow("SELECT subscriber FROM subscribers WHERE auth=?", authStr).Scan(&s) // work for?

	if err != nil {
		log.Print(err)
		return nil, err
	}
	return s, nil
}

func (sr *SubscriptionResource) QueryByAuthToken(query string) (bool, error) { // do the check each time??
	var existingToken string
	// guard for type this way what if the env didnt have teh right type!
	err := db.QueryRow(query).Scan(&existingToken)
	if err != nil {
		return false, err
	}
	return true, nil
}
