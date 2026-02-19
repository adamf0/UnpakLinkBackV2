package application

import (
	domainClick "UnpakSiamida/modules/click/domain"
	"context"
	"time"
)

type GetAllClicksQueryHandler struct {
	Repo domainClick.IClickRepository
}

func (h *GetAllClicksQueryHandler) Handle(
	ctx context.Context,
	q GetAllClicksQuery,
) (domainClick.PagedClicks, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	Clicks, total, err := h.Repo.GetAll(
		ctx,
		q.Search,
		q.SearchFilters,
		q.Page,
		q.Limit,
	)
	if err != nil {
		return domainClick.PagedClicks{}, err
	}

	currentPage := 1
	totalPages := 1

	if q.Page != nil {
		currentPage = *q.Page
	}
	if q.Limit != nil && *q.Limit > 0 {
		totalPages = int((total + int64(*q.Limit) - 1) / int64(*q.Limit))
	}

	return domainClick.PagedClicks{
		Data:        Clicks,
		Total:       total,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
	}, nil
}
