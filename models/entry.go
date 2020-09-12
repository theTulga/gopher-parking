package models

import (
	"fmt"
	"time"
)

type Entry struct {
	BillingAmount int
	EnteredAt     time.Time
	ExitedAt      time.Time
}

func (e Entry) String() string {
	return fmt.Sprintf("%d %s %s",
		e.EnteredAt.Format("2006-01-02 15:04:05"),
		e.ExitedAt.Format("2006-01-02 15:04:05"),
		e.BillingAmount)
}
