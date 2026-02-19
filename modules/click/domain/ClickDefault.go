package domain

import (
	"time"
)

type ClickDefault struct {
	LinkId    uint
	LinkShort string
	LinkLong  string
	IP        string
	IpClient  string
	ISO       string
	Country   string
	UserAgent string
	Created   time.Time
}
