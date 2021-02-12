package adminapi

import (
	"github.com/danieeelfr/myscheduler/internal/config"
	"github.com/danieeelfr/myscheduler/pkg/wait"
	"github.com/labstack/echo/v4"
)

type handler struct {
	wait *wait.Wait
}

func newHandler(cfg *config.Config, wg *wait.Wait) (*handler, error) {

	h := new(handler)
	h.wait = wg

	return h, nil
}

func (h *handler) GetUsers(ctx echo.Context) error {

	return nil
}
