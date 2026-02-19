package application

import "UnpakSiamida/common/domain"

type GetAllClicksQuery struct {
	Search        string
	SearchFilters []domain.SearchFilter
	Page          *int
	Limit         *int
}
