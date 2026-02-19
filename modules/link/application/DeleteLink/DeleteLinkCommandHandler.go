package application

import (
	"context"

	domainlink "UnpakSiamida/modules/link/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeleteLinkCommandHandler struct {
	Repo domainlink.ILinkRepository
}

func (h *DeleteLinkCommandHandler) Handle(
	ctx context.Context,
	cmd DeleteLinkCommand,
) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Validate UUID
	linkUUID, err := uuid.Parse(cmd.Uuid)
	if err != nil {
		return "", domainlink.InvalidUuid()
	}

	// Get existing link
	prev, err := h.Repo.GetByUuid(ctx, linkUUID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", domainlink.NotFound(cmd.Uuid)
		}
		return "", err
	}

	if prev.Creator != cmd.Creator {
		return "", domainlink.RejectDelete()
	}

	// Delete by UUID
	if err := h.Repo.Delete(ctx, linkUUID); err != nil { // FIXED
		return "", err
	}

	return cmd.Uuid, nil
}
