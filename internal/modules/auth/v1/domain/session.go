package domain

import "time"

type RefreshSession struct {
	JTI       string
	UserID    uint
	ExpiresAt time.Time
	Revoked   bool
}
