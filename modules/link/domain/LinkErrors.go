package domain

import (
	"UnpakSiamida/common/domain"
	"fmt"
)

func EmptyData() domain.Error {
	return domain.NotFoundError("Link.EmptyData", "data is not found")
}

func InvalidUuid() domain.Error {
	return domain.NotFoundError("Link.InvalidUuid", "uuid is invalid")
}

func InvalidState() domain.Error {
	return domain.NotFoundError("Link.InvalidState", "state is invalid")
}

func InvalidData() domain.Error {
	return domain.NotFoundError("Link.InvalidData", "data is invalid")
}

func NotUnique() domain.Error {
	return domain.NotFoundError("Link.NotUnique", "link not unique")
}
func InvalidFormatStart() domain.Error {
	return domain.NotFoundError("Link.InvalidFormatStart", "date time start is invalid format")
}
func InvalidFormatEnd() domain.Error {
	return domain.NotFoundError("Link.InvalidFormatEnd", "date time end is invalid format")
}
func RejectDelete() domain.Error {
	return domain.NotFoundError("Link.RejectDelete", "delete shortlink is rejected because different creator")
}

func NotFound(id string) domain.Error {
	return domain.NotFoundError("Link.NotFound", fmt.Sprintf("Link with identifier %s not found", id))
}
