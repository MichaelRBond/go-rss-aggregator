package types

import (
	"database/sql"
	"net/http"

	"github.com/michaelrbond/go-rss-aggregator/configuration"
)

// Context holds context that the controllers may need
type Context struct {
	Config configuration.Config
	Db     *sql.DB
}

// HTTPHandler is a generic handler
type HTTPHandler struct {
	*Context
	H func(http.ResponseWriter, *http.Request, *Context)
}

func (h HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.H(res, req, h.Context)
}