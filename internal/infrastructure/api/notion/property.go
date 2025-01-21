package notion

import (
	"errors"
	"strings"
	"time"

	"github.com/jomei/notionapi"
)

// PropertyToDatetime converts the value of given property to date
func PropertyToDatetime(prop notionapi.Property) (time.Time, error) {
	timeStr, err := PropertyToString(prop)
	if err != nil {
		return time.Time{}, err
	}

	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, errors.Join(errors.New("unable parse property to time"), err)
	}
	return parsedTime, nil
}

// PropertyToString convert the value of given property to string
func PropertyToString(prop notionapi.Property) (string, error) {
	t := string(prop.GetType())
	switch {
	case t == "rich_text":
		if p, ok := prop.(*notionapi.RichTextProperty); ok {
			return fromRichTextToString(p.RichText), nil
		}
	case t == "date":
		if p, ok := prop.(*notionapi.DateProperty); ok {
			return p.Date.Start.String(), nil
		}
	case t == "title":
		if p, ok := prop.(*notionapi.TitleProperty); ok {
			return fromRichTextToString(p.Title), nil
		}
	default:
		return "", errors.New("undefined property type: " + t)
	}

	return "", errors.New("unable to convert property with type: " + t)
}

// fromRichTextToString converts notion rich-text property to string
func fromRichTextToString(rt []notionapi.RichText) string {
	out := ""
	for _, block := range rt {
		out = out + " " + block.PlainText
	}
	return strings.TrimSpace(out)
}
