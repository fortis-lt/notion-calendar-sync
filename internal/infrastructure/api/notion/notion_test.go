package notion

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"fortis.notion-calendar-sync/internal/config"
	"github.com/jomei/notionapi"
	"github.com/onsi/gomega"
)

func TestNotionCli(t *testing.T) {
	g := gomega.NewWithT(t)

	// notion client
	n := testNotionCli(t)

	// events check
	events, err := n.Events(context.TODO())
	g.Expect(err).To(gomega.BeNil())
	fmt.Println(events)
}

func TestNotionEventsFilter(t *testing.T) {
	g := gomega.NewWithT(t)

	// notion client
	n := testNotionCli(t)

	flt, err := n.eventsFilter()
	g.Expect(err).To(gomega.BeNil())
	g.Expect(flt.Filter).ToNot(gomega.BeNil())
}

func testNotionCli(t *testing.T) *Notion {
	g := gomega.NewWithT(t)

	rootDir, _ := os.Getwd()
	cnf, err := config.New(filepath.Join(rootDir, "../../../../config/config.yaml"))
	g.Expect(err).To(gomega.BeNil())

	return &Notion{
		client: notionapi.NewClient(notionapi.Token(cnf.Infrastructure.Notion.IntegrationKey)),
		config: cnf.Infrastructure.Notion,
	}
}
