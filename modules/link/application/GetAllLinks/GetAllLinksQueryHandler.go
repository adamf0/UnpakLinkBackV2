package application

import (
	domainLink "UnpakSiamida/modules/link/domain"
	"context"
	"time"
)

type GetAllLinksQueryHandler struct {
	Repo domainLink.ILinkRepository
}

func (h *GetAllLinksQueryHandler) Handle(
	ctx context.Context,
	q GetAllLinksQuery,
) (domainLink.PagedLinks, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	Links, total, err := h.Repo.GetAll(
		ctx,
		q.Search,
		q.SearchFilters,
		q.Page,
		q.Limit,
	)
	if err != nil {
		return domainLink.PagedLinks{}, err
	}

	currentPage := 1
	totalPages := 1

	if q.Page != nil {
		currentPage = *q.Page
	}
	if q.Limit != nil && *q.Limit > 0 {
		totalPages = int((total + int64(*q.Limit) - 1) / int64(*q.Limit))
	}

	return domainLink.PagedLinks{
		Data:        Links,
		Total:       total,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
	}, nil
}
