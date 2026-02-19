package application

import (
	"context"
	"errors"
	"os"

	helper "UnpakSiamida/common/helper"
	domainaccount "UnpakSiamida/modules/account/domain"
	"time"

	"gorm.io/gorm"
)

type LoginCommandHandler struct {
	Repo domainaccount.IAccountRepository
}

func (h *LoginCommandHandler) Handle(
	ctx context.Context,
	cmd LoginCommand,
) (*domainaccount.LoginResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	sid := ""
	if cmd.Username == os.Getenv("USERNAME_AUTH") && cmd.Password == os.Getenv("PASSWORD_AUTH") {
		sid = "putiklink"
	} else {
		user, err := h.Repo.Auth(ctx, cmd.Username, cmd.Password)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, domainaccount.InvalidCredential()
			}
			return nil, err
		}

		sid = user.ID
	}

	accessToken, refreshToken, err := helper.GenerateToken(sid)
	if err != nil {
		return nil, err
	}

	return &domainaccount.LoginResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       sid,
	}, nil
}
