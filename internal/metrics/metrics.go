package metrics

import (
	"expvar"
	_ "expvar"
)

type collector struct {
	reqCount      *expvar.Int
	docsGenerated *expvar.Int
	errCount      *expvar.Int
}

type Collector interface {
	IncReq()
	IncDocs()
	IncErrs()
}

func (c *collector) IncReq() {
	c.reqCount.Add(1)
}

func (c *collector) IncDocs() {
	c.docsGenerated.Add(1)
}

func (c *collector) IncErrs() {
	c.errCount.Add(1)
}

func NewMetricsCollector() Collector {
	c := collector{
		reqCount:      expvar.NewInt("req_count"),
		docsGenerated: expvar.NewInt("docs_generated"),
		errCount:      expvar.NewInt("err_count"),
	}
	return &c
}
