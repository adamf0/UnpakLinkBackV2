package application

import (
	"context"

	domainlink "UnpakSiamida/modules/link/domain"
	"time"
)

type CreateLinkCommandHandler struct {
	Repo domainlink.ILinkRepository
}

func (h *CreateLinkCommandHandler) Handle(
	ctx context.Context,
	cmd CreateLinkCommand,
) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	countLink := h.Repo.CountLink(ctx, cmd.ShortUrl, cmd.Creator)

	result := domainlink.NewLink(
		cmd.ShortUrl,
		cmd.LongUrl,
		cmd.Password,
		cmd.Start,
		cmd.End,
		cmd.Creator,
		countLink == 0,
	)

	if !result.IsSuccess {
		return "", result.Error
	}

	createLink := result.Value
	if err := h.Repo.Create(ctx, createLink); err != nil {
		return "", err
	}

	return result.Value.UUID.String(), nil
}
