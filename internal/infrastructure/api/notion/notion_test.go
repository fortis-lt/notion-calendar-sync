package notion

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"fortis.notion-calendar-sync/internal/config"
	"fortis.notion-calendar-sync/internal/domain"
	"github.com/jomei/notionapi"
	"github.com/onsi/gomega"
)

func TestNotionCli(t *testing.T) {
	g := gomega.NewWithT(t)

	// notion client
	n := testNotionCli()

	// events check
	events, err := n.Events(context.TODO())
	g.Expect(err).To(gomega.BeNil())
	fmt.Println(events)
}

func TestNotionEventsFilter(t *testing.T) {
	g := gomega.NewWithT(t)

	// notion client
	n := testNotionCli()

	flt, err := n.eventsFilter()
	g.Expect(err).To(gomega.BeNil())
	g.Expect(flt.Filter).ToNot(gomega.BeNil())
}

func TestHandleResponsePage(t *testing.T) {
	g := gomega.NewWithT(t)

	// notion client
	n := testNotionCli()

	tests := []struct {
		description string
		page        notionapi.Page
		res         *domain.NotionEvent
		err         error
	}{
		{
			description: "valid",
			page: notionapi.Page{
				ID:  "dolores",
				URL: "https://hard-spiderling.name",
				Properties: notionapi.Properties{
					"Name": &notionapi.TitleProperty{
						ID:   "doloremque",
						Type: "title",
						Title: []notionapi.RichText{
							{PlainText: "name of the page"},
						},
					},
					"Event-Id": &notionapi.RichTextProperty{
						ID:   "nihil",
						Type: "rich_text",
						RichText: []notionapi.RichText{
							{PlainText: "event-id"},
						},
					},
					"Date": &notionapi.DateProperty{
						ID:   "est",
						Type: "date",
						Date: &notionapi.DateObject{
							Start: &notionapi.Date{},
						},
					},
				},
			},
			res: &domain.NotionEvent{
				Id:       "dolores",
				Url:      "https://hard-spiderling.name",
				Name:     "name of the page",
				RefId:    "event-id",
				Datetime: time.Time{},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			event, err := n.handleResponsePage(tt.page)
			if tt.err != nil {
				// TODO: deeper reserc needed. errors.Is(err, tt.err) doesn't work properly
				g.Expect(err).ToNot(gomega.BeNil())
			} else {
				g.Expect(err).To(gomega.BeNil())
				g.Expect(event).To(gomega.Equal(tt.res))
			}
		})
	}
}

func testNotionCli() *Notion {
	cnf := config.AppConfig{
		Global: &config.AppGlobal{
			LogLevel: "debug",
		},
		Infrastructure: &config.AppInfra{
			Notion: &config.NotionConfig{
				IntegrationKey: os.Getenv("TEST_NOTION_INTEGRATION_KEY"),
				Database: &config.NotionDatabaseConfig{
					Id:     "e853a9b1eb8441cba2815f90949ec0ff",
					Name:   "tasks",
					Filter: "TBD",
					Properties: &config.NotionDatabasePropertiesConfig{
						Name:     "Name",
						RefId:    "Event-Id",
						Datetime: "Date",
					},
				},
			},
		},
	}

	return &Notion{
		client: notionapi.NewClient(notionapi.Token(cnf.Infrastructure.Notion.IntegrationKey)),
		config: cnf.Infrastructure.Notion,
	}
}
