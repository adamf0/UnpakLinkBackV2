package application

import (
	"context"

	domainlink "UnpakSiamida/modules/link/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RollbackPasswordCommandHandler struct {
	Repo domainlink.ILinkRepository
}

func (h *RollbackPasswordCommandHandler) Handle(
	ctx context.Context,
	cmd RollbackPasswordCommand,
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
	result := domainlink.RollbackPassword(
		existingLink,
		linkUUID,
		cmd.Creator,
	)

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
