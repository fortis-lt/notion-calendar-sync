package notion

import (
	"errors"
	"testing"
	"time"

	"github.com/jomei/notionapi"
	"github.com/onsi/gomega"
)

func TestPropertyToDatetime(t *testing.T) {
	g := gomega.NewWithT(t)

	tests := []struct {
		description string
		property    notionapi.Property
		res         time.Time
		err         error
	}{
		{
			description: "date",
			property: &notionapi.DateProperty{
				ID:   "est",
				Type: "date",
				Date: &notionapi.DateObject{
					Start: &notionapi.Date{},
				},
			},
			res: time.Time{},
			err: nil,
		},
		{
			description: "rich_text",
			property: &notionapi.RichTextProperty{
				ID:   "nihil",
				Type: "rich_text",
				RichText: []notionapi.RichText{
					{PlainText: "ipsa vel cum"},
				},
			},
			res: time.Time{},
			err: errors.New("unable parse property to time"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			res, err := PropertyToDatetime(tt.property)
			if tt.err != nil {
				// TODO: deeper reserc needed. errors.Is(err, tt.err) doesn't work properly
				g.Expect(err).ToNot(gomega.BeNil())
			} else {
				g.Expect(err).To(gomega.BeNil())
				g.Expect(res).To(gomega.Equal(tt.res))
			}
		})
	}
}

func TestPropertyToString(t *testing.T) {
	g := gomega.NewWithT(t)

	tests := []struct {
		description string
		property    notionapi.Property
		res         string
		err         error
	}{
		{
			description: "rich_text",
			property: &notionapi.RichTextProperty{
				ID:   "nihil",
				Type: "rich_text",
				RichText: []notionapi.RichText{
					{PlainText: "ipsa vel cum"},
					{PlainText: "iest quo perspiciatis"},
				},
			},
			res: "ipsa vel cum iest quo perspiciatis",
			err: nil,
		},
		{
			description: "title",
			property: &notionapi.TitleProperty{
				ID:   "doloremque",
				Type: "title",
				Title: []notionapi.RichText{
					{PlainText: "ipsa vel cum"},
					{PlainText: "iest quo perspiciatis"},
				},
			},
			res: "ipsa vel cum iest quo perspiciatis",
			err: nil,
		},
		{
			description: "date",
			property: &notionapi.DateProperty{
				ID:   "est",
				Type: "date",
				Date: &notionapi.DateObject{
					Start: &notionapi.Date{},
				},
			},
			res: "0001-01-01T00:00:00Z",
			err: nil,
		},
		{
			description: "number",
			property: &notionapi.NumberProperty{
				ID:     "reiciendis",
				Type:   "number",
				Number: 1,
			},
			res: "",
			err: errors.New("undefined property type: number"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			res, err := PropertyToString(tt.property)
			if tt.err != nil {
				g.Expect(err).To(gomega.Equal(tt.err))
			} else {
				g.Expect(err).To(gomega.BeNil())
				g.Expect(res).To(gomega.Equal(tt.res))
			}
		})
	}
}
