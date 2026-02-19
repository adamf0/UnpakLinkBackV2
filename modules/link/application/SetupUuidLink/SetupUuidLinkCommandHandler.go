package application

import (
	"context"

	domainlink "UnpakSiamida/modules/link/domain"
)

type SetupUuidLinkCommandHandler struct {
	Repo domainlink.ILinkRepository
}

func (h *SetupUuidLinkCommandHandler) Handle(
	ctx context.Context,
	cmd SetupUuidLinkCommand,
) (string, error) {

	err := h.Repo.SetupUuid(ctx)
	if err != nil {
		return "", err
	}

	return "berhasil setup uuid pada data", nil
}
