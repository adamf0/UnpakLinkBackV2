package application

import (
	"context"

	commondomain "UnpakSiamida/common/domain"
	domainlink "UnpakSiamida/modules/link/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MoveLinkCommandHandler struct {
	Repo domainlink.ILinkRepository
}

func (h *MoveLinkCommandHandler) Handle(
	ctx context.Context,
	cmd MoveLinkCommand,
) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// -------------------------
	// VALIDATE UUID
	// -------------------------
	linkUUID, err := uuid.Parse(cmd.Uuid)
	if err != nil {
		return "", domainlink.InvalidUuid()
	}

	// -------------------------
	// GET EXISTING link
	// -------------------------
	existingLink, err := h.Repo.GetByUuid(ctx, linkUUID) // ← memastikan pakai nama interface yg benar
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", domainlink.NotFound(cmd.Uuid)
		}
		return "", err
	}

	// -------------------------
	// AGGREGATE ROOT LOGIC
	// -------------------------
	var result commondomain.ResultValue[*domainlink.Link]

	switch cmd.State {
	case "delete":
		result = domainlink.MoveDelete(existingLink, linkUUID, cmd.Creator)

	case "archive":
		result = domainlink.MoveArchive(existingLink, linkUUID, cmd.Creator)

	case "active":
		result = domainlink.MoveActive(existingLink, linkUUID, cmd.Creator)

	default:
		return "", domainlink.InvalidState()
	}

	if !result.IsSuccess {
		return "", result.Error
	}

	updatedLink := result.Value

	// -------------------------
	// SAVE TO REPOSITORY
	// -------------------------
	if err := h.Repo.Update(ctx, updatedLink); err != nil {
		return "", err
	}

	return updatedLink.UUID.String(), nil
}
