package model

import "database/sql"

type NotificationResource struct {
	*sql.DB
}

func NewNotificationResource(db *sql.DB) (*DB, error) {
	return &DB{db}, nil
}

type Notification struct {
	Id          int32  `json:"id"`
	Created     int32  `json:"created"`
	Status      int32  `json:"status"`
	Title       string `json:"title"` // binding req??
	Description string `json:"description"`
}

// something like this, note this is a placeholder
const (
	NotificationStatus = "notify"
	SeenStatus         = "seen"
)

func (nr *NotificationResource) GetNotifications(s Subscription) error {

	return db.QueryRow(sql)
}
