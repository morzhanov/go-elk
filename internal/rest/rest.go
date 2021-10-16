package rest

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/morzhanov/go-elk-example/internal/es"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type rest struct {
	esearch es.ElasticSearch
	log     *zap.Logger
}

type REST interface {
}

func (r *rest) handleHttpErr(c *gin.Context, err error) {
	c.String(http.StatusInternalServerError, err.Error())
	r.log.Error("error in the handler", zap.Error(err))
}

func (r *rest) handleFind(c *gin.Context) {
	fName := c.Param("field")
	val := c.Param("value")
	res, err := r.esearch.Find(fName, val)
	if err != nil {
		r.handleHttpErr(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (r *rest) handleUpdate(c *gin.Context) {
	// TODO: update
	//fName := c.Param("field")
	//val := c.Param("value")
	//res, err := r.esearch.Find(fName, val)
	//if err != nil {
	//	r.handleHttpErr(c, err)
	//	return
	//}
	//c.JSON(http.StatusOK, res)
}

func (r *rest) handleDelete(c *gin.Context) {
	// TODO: delete
	//fName := c.Param("field")
	//val := c.Param("value")
	//res, err := r.esearch.Find(fName, val)
	//if err != nil {
	//	r.handleHttpErr(c, err)
	//	return
	//}
	//c.JSON(http.StatusOK, res)
}

func (r *rest) Listen() error {
	router := gin.Default()
	config := cors.DefaultConfig()
	router.Use(cors.New(config))
	router.GET("/:field/:value", r.handleFind)
	router.PUT("/", r.handleUpdate)
	router.DELETE("/", r.handleDelete)
	return router.Run()
}

func NewREST(esearch es.ElasticSearch, l *zap.Logger) REST {
	return &rest{esearch, l}
}
