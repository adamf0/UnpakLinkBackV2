package domaintest

import (
	"testing"
	"time"

	domain "UnpakSiamida/modules/link/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ====================
// CREATE SUCCESS
// ====================
func TestNewLink_Success(t *testing.T) {
	short := "abc123"
	long := "https://google.com"
	creator := "admin"

	res := domain.NewLink(short, long, nil, nil, nil, creator, true)

	require.True(t, res.IsSuccess)
	link := res.Value
	require.NotNil(t, link)

	assert.Equal(t, short, link.ShortUrl)
	assert.Equal(t, long, link.LongUrl)
	assert.Equal(t, creator, link.Creator)
	assert.NotEqual(t, uuid.Nil, link.UUID)

	// Status default sekarang "active"
	require.NotNil(t, link.Status)
	assert.Equal(t, "active", *link.Status)
}

// ====================
// CREATE FAIL
// ====================
func TestNewLink_NotUnique(t *testing.T) {
	res := domain.NewLink("a", "b", nil, nil, nil, "admin", false)

	require.False(t, res.IsSuccess)
	assert.Equal(t, domain.NotUnique().Code, res.Error.Code)
}

// ====================
// UPDATE SUCCESS
// ====================
func TestUpdateLink_Success(t *testing.T) {
	res := domain.NewLink("abc", "long", nil, nil, nil, "admin", true)
	require.True(t, res.IsSuccess)

	prev := res.Value
	newShort := "updated"

	updateRes := domain.Update(prev, prev.UUID, newShort, "admin", true)

	require.True(t, updateRes.IsSuccess)
	assert.Equal(t, newShort, updateRes.Value.ShortUrl)
}

// ====================
// UPDATE FAIL CASES
// ====================
func TestUpdateLink_Fail(t *testing.T) {
	res := domain.NewLink("abc", "long", nil, nil, nil, "admin", true)
	require.True(t, res.IsSuccess)

	prev := res.Value

	tests := []struct {
		name     string
		prev     *domain.Link
		uid      uuid.UUID
		creator  string
		isUnique bool
		wantErr  string
	}{
		{
			name:     "PrevNil",
			prev:     nil,
			uid:      uuid.New(),
			creator:  "admin",
			isUnique: true,
			wantErr:  domain.EmptyData().Description,
		},
		{
			name:     "UUIDMismatch",
			prev:     prev,
			uid:      uuid.New(),
			creator:  "admin",
			isUnique: true,
			wantErr:  domain.InvalidData().Description,
		},
		{
			name:     "CreatorMismatch",
			prev:     prev,
			uid:      prev.UUID,
			creator:  "user",
			isUnique: true,
			wantErr:  domain.InvalidData().Description,
		},
		{
			name:     "NotUnique",
			prev:     prev,
			uid:      prev.UUID,
			creator:  "admin",
			isUnique: false,
			wantErr:  domain.NotUnique().Description,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := domain.Update(tt.prev, tt.uid, "new", tt.creator, tt.isUnique)

			require.False(t, res.IsSuccess)
			assert.Equal(t, tt.wantErr, res.Error.Description)
		})
	}
}

// ====================
// STATUS MOVEMENT TEST
// ====================
func TestMoveStatus_Success(t *testing.T) {
	res := domain.NewLink("a", "b", nil, nil, nil, "admin", true)
	require.True(t, res.IsSuccess)

	prev := res.Value

	// Default status harus active
	require.NotNil(t, prev.Status)
	assert.Equal(t, "active", *prev.Status)

	// Move Delete
	delRes := domain.MoveDelete(prev, prev.UUID, "admin")
	require.True(t, delRes.IsSuccess)
	require.NotNil(t, delRes.Value.Status)
	assert.Equal(t, "delete", *delRes.Value.Status)

	// Move Archive
	arcRes := domain.MoveArchive(prev, prev.UUID, "admin")
	require.True(t, arcRes.IsSuccess)
	require.NotNil(t, arcRes.Value.Status)
	assert.Equal(t, "archive", *arcRes.Value.Status)

	// Move Active (kembali aktif)
	actRes := domain.MoveActive(prev, prev.UUID, "admin")
	require.True(t, actRes.IsSuccess)
	assert.Nil(t, actRes.Value.Status)
}

// ====================
// GIVE TIME SUCCESS
// ====================
func TestGiveTime_Success(t *testing.T) {
	res := domain.NewLink("a", "b", nil, nil, nil, "admin", true)
	require.True(t, res.IsSuccess)

	prev := res.Value

	start := time.Now().Format("2006-01-02 15:04:05")
	end := time.Now().Add(time.Hour).Format("2006-01-02 15:04:05")

	timeRes := domain.GiveTime(prev, prev.UUID, start, end, "admin")

	require.True(t, timeRes.IsSuccess)
	assert.NotNil(t, timeRes.Value.StartAccess)
	assert.NotNil(t, timeRes.Value.EndAccess)
}

// ====================
// GIVE TIME FAIL FORMAT
// ====================
func TestGiveTime_InvalidFormat(t *testing.T) {
	res := domain.NewLink("a", "b", nil, nil, nil, "admin", true)
	require.True(t, res.IsSuccess)

	prev := res.Value

	timeRes := domain.GiveTime(prev, prev.UUID, "salah-format", "2024-01-01", "admin")

	require.False(t, timeRes.IsSuccess)
	assert.Equal(t, domain.InvalidFormatStart().Code, timeRes.Error.Code)
}

// ====================
// PASSWORD TEST
// ====================
func TestPassword_Success(t *testing.T) {
	res := domain.NewLink("a", "b", nil, nil, nil, "admin", true)
	require.True(t, res.IsSuccess)

	prev := res.Value

	passRes := domain.GivePassword(prev, prev.UUID, "secret", "admin")
	require.True(t, passRes.IsSuccess)
	assert.Equal(t, "secret", *passRes.Value.Password)

	rollback := domain.RollbackPassword(prev, prev.UUID, "admin")
	require.True(t, rollback.IsSuccess)
	assert.Nil(t, rollback.Value.Password)
}

// ====================
// ROLLBACK TIME SUCCESS
// ====================
func TestRollbackTime_Success(t *testing.T) {
	now := time.Now()
	l := &domain.Link{
		UUID:        uuid.New(),
		Creator:     "adam",
		StartAccess: &now,
		EndAccess:   &now,
	}

	res := domain.RollbackTime(l, l.UUID, "adam")

	require.True(t, res.IsSuccess)
	assert.Nil(t, res.Value.StartAccess)
	assert.Nil(t, res.Value.EndAccess)
}

// ====================
// ROLLBACK TIME FAIL
// ====================
func TestRollbackTime_Fail(t *testing.T) {
	prev := &domain.Link{
		UUID:    uuid.New(),
		Creator: "adam",
	}

	tests := []struct {
		name    string
		prev    *domain.Link
		uid     uuid.UUID
		creator string
		wantErr string
	}{
		{
			name:    "PrevNil",
			prev:    nil,
			uid:     uuid.New(),
			creator: "adam",
			wantErr: domain.EmptyData().Description,
		},
		{
			name:    "UUIDMismatch",
			prev:    prev,
			uid:     uuid.New(),
			creator: "adam",
			wantErr: domain.InvalidData().Description,
		},
		{
			name:    "CreatorMismatch",
			prev:    prev,
			uid:     prev.UUID,
			creator: "hacker",
			wantErr: domain.InvalidData().Description,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := domain.RollbackTime(tt.prev, tt.uid, tt.creator)

			require.False(t, res.IsSuccess)
			assert.Equal(t, tt.wantErr, res.Error.Description)
		})
	}
}

// ====================
// ROLLBACK PASSWORD SUCCESS
// ====================
func TestRollbackPassword_Success(t *testing.T) {
	pass := "rahasia"

	l := &domain.Link{
		UUID:     uuid.New(),
		Creator:  "adam",
		Password: &pass,
	}

	res := domain.RollbackPassword(l, l.UUID, "adam")

	require.True(t, res.IsSuccess)
	assert.Nil(t, res.Value.Password)
}

// ====================
// ROLLBACK PASSWORD FAIL
// ====================
func TestRollbackPassword_Fail(t *testing.T) {
	prev := &domain.Link{
		UUID:    uuid.New(),
		Creator: "adam",
	}

	tests := []struct {
		name    string
		prev    *domain.Link
		uid     uuid.UUID
		creator string
		wantErr string
	}{
		{
			name:    "PrevNil",
			prev:    nil,
			uid:     uuid.New(),
			creator: "adam",
			wantErr: domain.EmptyData().Description,
		},
		{
			name:    "UUIDMismatch",
			prev:    prev,
			uid:     uuid.New(),
			creator: "adam",
			wantErr: domain.InvalidData().Description,
		},
		{
			name:    "CreatorMismatch",
			prev:    prev,
			uid:     prev.UUID,
			creator: "hacker",
			wantErr: domain.InvalidData().Description,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := domain.RollbackPassword(tt.prev, tt.uid, tt.creator)

			require.False(t, res.IsSuccess)
			assert.Equal(t, tt.wantErr, res.Error.Description)
		})
	}
}
