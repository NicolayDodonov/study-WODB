package graphQl

import (
	"net/http"
	"study-WODB/internal/config"
	"study-WODB/internal/logger"
)

type GraphQL struct {
}

func NewGraphQL(c *config.Config, l *logger.Logger) *GraphQL {
	return &GraphQL{}
}

func (ga *GraphQL) Check(w http.ResponseWriter, r *http.Request) {

}
