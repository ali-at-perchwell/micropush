package model

import (
	"database/sql"
	"fmt"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
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

func (sr *SubscriptionResource) Create(sub Subscription) error { // subscription created from json req
	_, err = sr.prepareDb(createSql).Exec(sub.Endpoint, sub.Keys.Auth, time.Now().Unix())

	if err != nil {
		return err
	}
	return nil
}

func (sr *SubscriptionResource) Delete(sub Subscription) error {
	_, err = sr.prepareDb(deleteSql).Exec(sub.Endpoint, sub.Keys.Auth, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (sr *SubscriptionResource) Update(sub Subscription) error {
	_, err = sr.prepareDb(updateSql).Exec(sub.Endpoint, sub.Keys.Auth, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func NewSubscriptionFromRequest(c *gin.Context) (Subscription, error) { // bind this????
	var s Subscription

	if err := c.BindJSON(&s); err != nil {
		return nil.error
	}

	var existingToken string
	exists, err := subscription.QueryByAuthToken(s.Keys.Auth)
	if err != nil {
		return nil, err
	}
	if exists {
		// TODO: ADD UPDATE
		// SUBSC TIME ETC

		// or if exists return true, return the subscription object so you dont have to query again?
		return subscription.Update(s.Keys.Auth)
	}

	// check to see if subscription exists in the DB

	c.String(200, "Success")
	return subscription // does this need a * or a &

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

func (sr *SubscriptionResource) getAuthToken(c *gin.Context) {

}
func (sr *SubscriptionResource) GetSubscription(c *gin.Context) {
	id, err := sr.getAuthToken(c)
}
