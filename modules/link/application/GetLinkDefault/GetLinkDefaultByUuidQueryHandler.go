package application

import (
	"context"

	domainLink "UnpakSiamida/modules/link/domain"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GetLinkDefaultByUuidQueryHandler struct {
	Repo domainLink.ILinkRepository
}

func (h *GetLinkDefaultByUuidQueryHandler) Handle(
	ctx context.Context,
	q GetLinkDefaultByUuidQuery,
) (*domainLink.LinkDefault, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	parsed, err := uuid.Parse(q.Uuid)
	if err != nil {
		return nil, domainLink.NotFound(q.Uuid)
	}

	Link, err := h.Repo.GetDefaultByUuid(ctx, parsed)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainLink.NotFound(q.Uuid)
		}
		return nil, err
	}

	return Link, nil
}
