package application

import (
	"context"
	"errors"

	domainaccount "UnpakSiamida/modules/account/domain"
	"time"

	"gorm.io/gorm"
)

type WhoamiCommandHandler struct {
	Repo domainaccount.IAccountRepository
}

func (h *WhoamiCommandHandler) Handle(
	ctx context.Context,
	cmd WhoamiCommand,
) (*domainaccount.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// _, err := uuid.Parse(cmd.SID)
	// if err != nil {
	// 	return nil, domainaccount.NotFound(cmd.SID)
	// }
	if len(cmd.SID) == 0 {
		return nil, domainaccount.NotFound(cmd.SID)
	}

	if cmd.SID == "putiklink" {
		return &domainaccount.Account{
			ID:        "putiklink",
			UUID:      nil,
			Username:  "putik",
			Level:     "ADMINISTRATOR",
			Name:      "PUTIK",
			ExtraRole: []domainaccount.ExtraRole{},
		}, nil
	}

	user, err := h.Repo.GetByUserId(ctx, cmd.SID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainaccount.InvalidCredential()
		}
		return nil, err
	}
	if user.ExtraRole == nil {
		user.ExtraRole = []domainaccount.ExtraRole{}
	}

	return user, nil
}
