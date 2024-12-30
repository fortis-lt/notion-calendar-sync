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

	rootDir, _ := os.Getwd()
	cnf, err := config.New(filepath.Join(rootDir, "../../../../config/config.yaml"))
	g.Expect(err).To(gomega.BeNil())

	n := &Notion{
		client: notionapi.NewClient(notionapi.Token(cnf.Infrastructure.Notion.IntegrationKey)),
		config: cnf.Infrastructure.Notion,
	}

	// events check
	events, err := n.Events(context.TODO())
	g.Expect(err).To(gomega.BeNil())
	fmt.Println(events)

}
