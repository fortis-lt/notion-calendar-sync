package notion

import (
	"fortis.notion-calendar-sync/internal/config"
	"fortis.notion-calendar-sync/internal/infrastructure"
	"github.com/jomei/notionapi"
)

type Notion struct {
	client *notionapi.Client
	config *config.NotionConfig
}

// New return a new Notion client based on provided configuration
func New(cnf *config.NotionConfig) infrastructure.Notion {
	return Notion{
		client: notionapi.NewClient(notionapi.Token(cnf.IntegrationKey)),
		config: cnf,
	}
}
