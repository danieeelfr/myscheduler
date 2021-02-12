package adminapi

import (
	"net/http"
	"path"

	"github.com/danieeelfr/myscheduler/internal/config"
	"github.com/danieeelfr/myscheduler/pkg/wait"
	"github.com/labstack/echo/v4"
)

const (
	Prefix       = "v1"
	UserEndpoint = "user"
)

type router struct {
	e      *echo.Echo
	routes []*route
}

type route struct {
	method   string
	endpoint string
	handler  func(c echo.Context) error
}

func newRouter(e *echo.Echo, cfg *config.Config, wg *wait.Wait) (*router, error) {

	handler, err := newHandler(cfg, wg)
	if err != nil {
		return nil, err
	}

	return &router{
		e: e,
		routes: []*route{
			&route{
				method:   http.MethodGet,
				endpoint: path.Join(Prefix, UserEndpoint),
				handler:  handler.GetUsers,
			},
		},
	}, nil
}

func (rtr *router) build() {

	for _, route := range rtr.routes {
		rtr.setRoute(route)
	}
}

func (rtr *router) setRoute(r *route) {
	switch r.method {
	case http.MethodGet:
		rtr.e.GET(r.endpoint, r.handler)
	case http.MethodPost:
		rtr.e.POST(r.endpoint, r.handler)
	case http.MethodPut:
		rtr.e.PUT(r.endpoint, r.handler)
	case http.MethodDelete:
		rtr.e.DELETE(r.endpoint, r.handler)
	default:
		log.Errorf("[%S] method not implemented", r.method)
	}
}
