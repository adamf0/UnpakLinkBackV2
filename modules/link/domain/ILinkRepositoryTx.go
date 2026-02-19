package domain

import (
	commoninfra "UnpakSiamida/common/infrastructure"
	"context"
)

type ILinkRepositoryTx interface {
	ILinkRepository
	InsertOutbox(ctx context.Context, msg *commoninfra.OutboxMessage) error
}
