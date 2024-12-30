package infrastructure

import (
	"context"

	"fortis.notion-calendar-sync/internal/domain"
)

type Notion interface {
	Events(ctx context.Context) ([]*domain.NotionEvent, error)
}
