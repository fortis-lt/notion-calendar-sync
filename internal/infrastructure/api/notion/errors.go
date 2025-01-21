package notion

import "errors"

var (
	ErrUndefinedPropertyType = errors.New("undefined property type")
)

const (
	errPropertyNotFound = "property '%s' not found"
	errorPropertyDecode = "property decoding error '%s'"
)
