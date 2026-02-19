package domain

type PagedClicks struct {
	Data        []ClickDefault `json:"data"`
	Total       int64          `json:"total"`
	CurrentPage int            `json:"current_page"`
	TotalPages  int            `json:"total_pages"`
}
