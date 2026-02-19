package application

import (
	"context"

	domainClick "UnpakSiamida/modules/click/domain"
	"time"
)

type CreateClickCommandHandler struct {
	Repo domainClick.IClickRepository
}

func (h *CreateClickCommandHandler) Handle(
	ctx context.Context,
	cmd CreateClickCommand,
) (uint, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result := domainClick.NewClick(
		cmd.LinkId,
		cmd.IP,
		cmd.IPClient,
		cmd.ISO,
		cmd.Country,
		cmd.Referer,
		cmd.UserAgent,
	)

	if !result.IsSuccess {
		return 0, result.Error
	}

	createClick := result.Value
	if err := h.Repo.Create(ctx, createClick); err != nil {
		return 0, err
	}

	return result.Value.ID, nil
}
