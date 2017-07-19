package model

type Subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     string `json:"keys"`
}

type Keys struct {
	P256DH string `json:"p256dh"`
	Auth   string `json:"auth"`
}

//Token  string `json:"token"`
//UserId string `json:"user_id"`
//BrowserClient??

type SubscriptionResource struct {
	db sql.DB // debating this
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

func GetSubscription(db *sql.DB) error {
	var subscription Subscription
	if c.Bind(&subscription) != nil {

	}
	sql := fmt.Sprintf("INSERT device_id FROM device_tokens WHERE device_id='%v'", deviceId)

}

func (s Subscription) Save(db *sql.DB) error { // or save??
	return nil
}
