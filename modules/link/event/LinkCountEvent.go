package event

import (
	"time"

	"github.com/google/uuid"
	// "UnpakSiamida/common/domain"
)

type LinkCountEvent struct {
	EventID    uuid.UUID
	LinkId     uint
	IP         string
	IpClient   string
	ISO        string
	Country    string
	Referer    string
	UserAgent  string
	OccurredOn time.Time
}

func NewLinkCountEvent(
	linkId uint,
	ip string,
	ipClient string,
	iso string,
	country string,
	referer string,
	userAgent string,
) *LinkCountEvent {

	return &LinkCountEvent{
		EventID:    uuid.New(),
		LinkId:     linkId,
		IP:         ip,
		IpClient:   ipClient,
		ISO:        iso,
		Country:    country,
		Referer:    referer,
		UserAgent:  userAgent,
		OccurredOn: time.Now().UTC(),
	}
}

func (e LinkCountEvent) ID() string {
	return e.EventID.String()
}

func (e LinkCountEvent) OccurredOnUTC() time.Time {
	return e.OccurredOn
}
