package infrastructure

import (
	domainLink "UnpakSiamida/modules/link/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type CrontabProcessor struct {
	DB *gorm.DB
}

func (p *CrontabProcessor) Process(ctx context.Context) error {
	var messages []domainLink.Link

	now := time.Now().UTC()

	if err := p.DB.
		Where("end_access > ?", now).
		Limit(10).
		Find(&messages).Error; err != nil {
		return err
	}

	for _, msg := range messages {
		p.DB.Model(&msg).Update("status", "archive").Update("updated_at", now)
	}

	return nil
}
