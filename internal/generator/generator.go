package generator

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/morzhanov/go-elk-example/internal/doc"
	"github.com/morzhanov/go-elk-example/internal/es"
	"go.uber.org/zap"
)

type gen struct {
	esearch es.ElasticSearch
	log     *zap.Logger
}

type Generator interface {
	Generate()
}

func (g *gen) generateDoc() (*doc.Document, error) {
	var d doc.Document
	if err := gofakeit.Struct(&d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (g *gen) Generate() {
	for {
		time.Sleep(time.Second * 30)
		d, err := g.generateDoc()
		if err != nil {
			g.log.Error("an error occurred when trying to generate doc", zap.Error(err))
			continue
		}
		if err := g.esearch.Save(d); err != nil {
			g.log.Error("an error occurred when trying to save doc to elasticsearch", zap.Error(err))
		}
	}
}

func NewGenerator(esearch es.ElasticSearch, l *zap.Logger) Generator {
	gofakeit.Seed(0)
	return &gen{esearch, l}
}
