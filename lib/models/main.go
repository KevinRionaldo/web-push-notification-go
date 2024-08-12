package models

import (
	"time"
)

type Push_notification struct {
	Id          string    `db:"id,omitempty"`
	Endpoint    string    `db:"endpoint"`
	Keys_p256dh string    `db:"keys_p256dh"`
	Keys_Auth   string    `db:"keys_auth"`
	Created_at  time.Time `db:"created_at,omitempty"`
	Updated_at  time.Time `db:"updated_at,omitempty"`
	Created_by  *string   `db:"created_by"`
	Updated_by  *string   `db:"updated_by"`
}
