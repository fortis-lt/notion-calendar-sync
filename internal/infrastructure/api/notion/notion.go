package notion

import (
	"context"
	"errors"
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
			event, err := n.handleResponsePage(p)
			if err != nil {
				return []*domain.NotionEvent{}, err
			}
			events = append(events, event)
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

// UpdateEvent updates the event
func (n *Notion) UpdateEvent(ctx context.Context, event *domain.NotionEvent) error {
	return errors.New("not implemented")
}

// handleResponsePage converts recieved notion page to internal notion event
func (n *Notion) handleResponsePage(p notionapi.Page) (*domain.NotionEvent, error) {
	out := &domain.NotionEvent{
		Id:  p.ID.String(),
		Url: p.URL,
	}

	// refId decoding
	prop, ok := p.Properties[n.config.Database.Properties.RefId]
	if !ok {
		return nil, fmt.Errorf(errPropertyNotFound, n.config.Database.Properties.RefId)
	}
	eventRefId, err := PropertyToString(prop)
	if err != nil {
		return nil, errors.Join(fmt.Errorf(errorPropertyDecode, "RefId"), err)
	}
	out.RefId = eventRefId

	// name decoding
	prop, ok = p.Properties[n.config.Database.Properties.Name]
	if !ok {
		return nil, fmt.Errorf(errPropertyNotFound, n.config.Database.Properties.Name)
	}
	eventName, err := PropertyToString(prop)
	if err != nil {
		return nil, errors.Join(fmt.Errorf(errorPropertyDecode, "Name"), err)
	}
	out.Name = eventName

	// date decoding
	prop, ok = p.Properties[n.config.Database.Properties.Datetime]
	if !ok {
		return nil, fmt.Errorf(errPropertyNotFound, n.config.Database.Properties.Datetime)
	}
	eventTime, err := PropertyToDatetime(prop)
	if err != nil {
		return nil, errors.Join(fmt.Errorf(errorPropertyDecode, "Datetime"), err)
	}
	out.Datetime = eventTime

	return out, nil

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
