package models

import (
	"github.com/softbrewery/gojoi/pkg/joi"
)

type Book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (bk *Book) ValidateSchema() error {
	schema := joi.Struct().Keys(joi.StructKeys{
		"ISBN":   joi.String().Regex(`^[0-9]{3}\-[0-9]{10}$`),
		"Title":  joi.String().NonZero().Required(),
		"Author": joi.String().NonZero().Required(),
	})
	if err := schema.Validate(bk); err != nil {
		return err
	}
	return nil
}
