package application

import "UnpakSiamida/common/domain"

type GetAllLinksQuery struct {
	Search        string
	SearchFilters []domain.SearchFilter
	Page          *int
	Limit         *int
}
