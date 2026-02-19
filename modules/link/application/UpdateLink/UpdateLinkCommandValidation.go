package application

import (
	helper "UnpakSiamida/common/helper"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func UpdateLinkCommandValidation(cmd UpdateLinkCommand) error {
	return validation.ValidateStruct(&cmd,
		validation.Field(&cmd.Uuid,
			validation.Required.Error("UUID cannot be blank"),
			validation.By(helper.ValidateUUIDv4),
			validation.By(helper.NoXSSFullScanWithDecode()),
		),

		validation.Field(&cmd.ShortUrl,
			validation.Required.Error("ShortUrl cannot be blank"),
			validation.By(helper.NoXSSFullScanWithDecode()),
		),
		validation.Field(&cmd.Creator,
			validation.Required.Error("Creator cannot be blank"),
			validation.By(helper.NoXSSFullScanWithDecode()),
		),
	)
}
