package domain

import (
	"UnpakSiamida/common/domain"
	"fmt"
)

func EmptyData() domain.Error {
	return domain.NotFoundError("Click.EmptyData", "data is not found")
}

func NotFound(id string) domain.Error {
	return domain.NotFoundError("Click.NotFound", fmt.Sprintf("Link with identifier %s not found", id))
}
