package document

import (
	"fmt"

	"github.com/gen2brain/go-fitz"
)

type Document struct {
	FitzDocument *fitz.Document
}

func NewDocument() (doc *Document, err error) {
	fitzDoc, err := fitz.NewFromMemory(nil)
	if err != nil {
		return nil, fmt.Errorf("fitz.NewFromMemory: %w", err)
	}

	doc = &Document{
		FitzDocument: fitzDoc,
	}

	return doc, nil
}
