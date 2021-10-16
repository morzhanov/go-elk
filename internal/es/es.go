package es

import "github.com/morzhanov/go-elk-example/internal/doc"

type es struct {
}

type ElasticSearch interface {
	Save(d *doc.Document) error
	Find(field string, value string) ([]*doc.Document, error)
	// TODO: investigate
	Update() error
	// TODO: investigate
	Delete() error
}

func NewES() ElasticSearch {
	return &es{}
}
