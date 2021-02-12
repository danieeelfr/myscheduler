package myscheduler

import (
	"github.com/danieeelfr/myscheduler/internal/config"
	"github.com/danieeelfr/myscheduler/internal/controller/server/adminapi"
	"github.com/danieeelfr/myscheduler/pkg/wait"
)

type Interactor interface {
	Start() error
	Shutdown()
}

func New(cfg *config.Config, wg *wait.Wait) (Interactor, error) {
	httpSrv, err := adminapi.New(cfg, wg)
	if err != nil {
		return nil, err
	}

	return httpSrv, nil
}
