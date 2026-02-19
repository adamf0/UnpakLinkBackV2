package domain

import (
	commonDomain "UnpakSiamida/common/domain"
	"context"

	"github.com/google/uuid"
)

type ILinkRepository interface {
	CountLink(ctx context.Context, link string, creator string) uint
	GetByUuid(ctx context.Context, uid uuid.UUID) (*Link, error)
	GetDefaultByShort(ctx context.Context, short string) (*LinkDefault, error)
	GetDefaultByUuid(ctx context.Context, uid uuid.UUID) (*LinkDefault, error)
	GetAll(
		ctx context.Context,
		search string,
		searchFilters []commonDomain.SearchFilter,
		page, limit *int,
	) ([]Link, int64, error)
	Create(ctx context.Context, jenisfile *Link) error
	Update(ctx context.Context, jenisfile *Link) error
	Delete(ctx context.Context, uid uuid.UUID) error
	SetupUuid(ctx context.Context) error
	WithTx(ctx context.Context, fn func(txRepo ILinkRepositoryTx) error) error
}
