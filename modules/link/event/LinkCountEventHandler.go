package event

import (
	"context"

	CreateClick "UnpakSiamida/modules/click/application/CreateClick"

	"github.com/mehdihadeli/go-mediatr"
	"gorm.io/gorm"
)

type LinkCountEventHandler struct {
	db *gorm.DB
}

func NewLinkCountEventHandler(
	db *gorm.DB,
) *LinkCountEventHandler {
	return &LinkCountEventHandler{
		db: db,
	}
}

func (h LinkCountEventHandler) Handle(
	ctx context.Context,
	event LinkCountEvent,
) error {
	cmd := CreateClick.CreateClickCommand{
		LinkId:    event.LinkId,
		IP:        event.IP,
		IPClient:  event.IpClient,
		ISO:       event.ISO,
		Country:   event.Country,
		Referer:   event.Referer,
		UserAgent: event.UserAgent,
	}

	_, err := mediatr.Send[CreateClick.CreateClickCommand, uint](context.Background(), cmd)
	if err != nil {
		return err
	}

	return nil
}
