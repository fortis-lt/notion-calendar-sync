package notion

import (
	"context"

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

// Events returns list of events based on the notion infra configuration
func (n *Notion) Events(ctx context.Context) ([]*domain.NotionEvent, error) {
	var events []*domain.NotionEvent

	filter, err := n.eventsFilter()
	if err != nil {
		return events, err
	}
	body := filter

	for {
		// request data from the notion
		resp, err := n.client.Database.Query(
			ctx,
			notionapi.DatabaseID(n.config.Database.Id),
			body,
		)
		if err != nil {
			return []*domain.NotionEvent{}, err
		}

		// collect recieved pages
		for _, p := range resp.Results {
			events = append(events, n.handleResponsePage(p))
		}

		if resp.HasMore && resp.NextCursor != "" {
			body = &notionapi.DatabaseQueryRequest{
				StartCursor: resp.NextCursor,
			}
		} else {
			break
		}

	}

	return events, nil
}

// handleResponsePage converts recieved notion page to internal notion event
func (n *Notion) handleResponsePage(p notionapi.Page) *domain.NotionEvent {
	event := domain.NotionEvent{
		Id: p.ID.String(),
	}
	return &event

}

// eventsFilter returns database efents filtration config
func (n *Notion) eventsFilter() (*notionapi.DatabaseQueryRequest, error) {
	r := &notionapi.DatabaseQueryRequest{}

	// TODO: the filer should be taken from config
	if n.config.Database.Filter == "" {
		return r, nil
	}

	r.Filter = notionapi.PropertyFilter{
		Property: "Status",
		Status: &notionapi.StatusFilterCondition{
			DoesNotEqual: "Done",
		},
	}

	return r, nil
}
