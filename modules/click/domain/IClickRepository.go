package domain

import (
	commondomain "UnpakSiamida/common/domain"
	"context"
)

type IClickRepository interface {
	Create(ctx context.Context, click *Click) error
	GetAll(
		ctx context.Context,
		search string,
		searchFilters []commondomain.SearchFilter,
		page, limit *int,
	) ([]ClickDefault, int64, error)
}
