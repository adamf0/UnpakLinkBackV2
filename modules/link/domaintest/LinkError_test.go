package domaintest

import (
	"testing"

	common "UnpakSiamida/common/domain"
	domain "UnpakSiamida/modules/link/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinkErrors(t *testing.T) {
	tests := []struct {
		name         string
		err          common.Error
		expectedCode string
		expectedDesc string
	}{
		{
			name:         "EmptyData_ReturnsCorrectError",
			err:          domain.EmptyData(),
			expectedCode: "Link.EmptyData",
			expectedDesc: "data is not found",
		},
		{
			name:         "InvalidUuid_ReturnsCorrectError",
			err:          domain.InvalidUuid(),
			expectedCode: "Link.InvalidUuid",
			expectedDesc: "uuid is invalid",
		},
		{
			name:         "InvalidState_ReturnsCorrectError",
			err:          domain.InvalidState(),
			expectedCode: "Link.InvalidState",
			expectedDesc: "state is invalid",
		},
		{
			name:         "InvalidData_ReturnsCorrectError",
			err:          domain.InvalidData(),
			expectedCode: "Link.InvalidData",
			expectedDesc: "data is invalid",
		},
		{
			name:         "NotUnique_ReturnsCorrectError",
			err:          domain.NotUnique(),
			expectedCode: "Link.NotUnique",
			expectedDesc: "link not unique",
		},
		{
			name:         "InvalidFormatStart_ReturnsCorrectError",
			err:          domain.InvalidFormatStart(),
			expectedCode: "Link.InvalidFormatStart",
			expectedDesc: "date time start is invalid format",
		},
		{
			name:         "InvalidFormatEnd_ReturnsCorrectError",
			err:          domain.InvalidFormatEnd(),
			expectedCode: "Link.InvalidFormatEnd",
			expectedDesc: "date time end is invalid format",
		},
		{
			name:         "RejectDelete_ReturnsCorrectError",
			err:          domain.RejectDelete(),
			expectedCode: "Link.RejectDelete",
			expectedDesc: "delete shortlink is rejected because different creator",
		},
		{
			name:         "NotFound_WithDynamicId_ReturnsCorrectError",
			err:          domain.NotFound("XYZ99"),
			expectedCode: "Link.NotFound",
			expectedDesc: "Link with identifier XYZ99 not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NotNil(t, tt.err)
			assert.Equal(t, tt.expectedCode, tt.err.Code)
			assert.Equal(t, tt.expectedDesc, tt.err.Description)
		})
	}
}
