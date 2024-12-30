package notion

import (
	"context"
	"fmt"

	"fortis.notion-calendar-sync/internal/config"
	"fortis.notion-calendar-sync/internal/domain"
	"fortis.notion-calendar-sync/internal/infrastructure"
	"github.com/jomei/notionapi"
)

type Notion struct {
	client *notionapi.Client
	config *config.NotionConfig
}

// New return a new Notion client based on provided configuration
func New(cnf *config.NotionConfig) infrastructure.Notion {
	return &Notion{
		client: notionapi.NewClient(notionapi.Token(cnf.IntegrationKey)),
		config: cnf,
	}
}

func (n *Notion) Events(ctx context.Context) ([]*domain.NotionEvent, error) {
	var events []*domain.NotionEvent

	// TODO: add filtration for the events querieng
	resp, err := n.client.Database.Query(ctx, notionapi.DatabaseID(n.config.Database.Id), nil)
	if err != nil {
		return events, err
	}

	fmt.Println(resp)

	return events, nil
}
