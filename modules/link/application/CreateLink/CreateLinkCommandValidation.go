package application

import (
	helper "UnpakSiamida/common/helper"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func CreateLinkCommandValidation(cmd CreateLinkCommand) error {
	return validation.ValidateStruct(&cmd,
		validation.Field(&cmd.ShortUrl,
			validation.Required.Error("ShortUrl cannot be blank"),
			validation.By(helper.NoXSSFullScanWithDecode()),
		),
		validation.Field(&cmd.LongUrl,
			validation.Required.Error("LongUrl cannot be blank"),
			validation.By(helper.NoXSSFullScanWithDecode()),
		),
		validation.Field(&cmd.Creator,
			validation.Required.Error("Creator cannot be blank"),
			validation.By(helper.NoXSSFullScanWithDecode()),
		),
	)
}
