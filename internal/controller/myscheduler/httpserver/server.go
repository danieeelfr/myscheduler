package httpserver

import (
	"context"

	"github.com/danieeelfr/myscheduler/internal/config"
	"github.com/danieeelfr/myscheduler/pkg/wait"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithField("package", "controller.httpserver")
)

type server struct {
	conf   *config.MySchedulerConf
	e      *echo.Echo
	router *router
	wait   *wait.Wait
}

func New(cfg *config.Config, wg *wait.Wait) (*server, error) {
	srv := new(server)
	srv.conf = cfg.GetMySchedulerConf()
	srv.wait = wg
	srv.e = echo.New()

	var err error
	srv.router, err = newRouter(srv.e, cfg, wg)
	if err != nil {
		return nil, err
	}

	return srv, nil

}

func (srv *server) Start() error {
	log.Infof("Starting http server. Host: [%s]", srv.conf.HttpServerHost)
	srv.wait.Add()

	srv.e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	srv.e.Use(middleware.Recover())

	srv.router.build()

	go func() {
		if err := srv.e.Start(srv.conf.HttpServerHost); err != nil {
			if !srv.wait.IsBlock() {
				log.WithError(err).Fatal("error starting http server. Error: [%s]", err)
			}

		}
	}()
	return nil
}

func (srv *server) Shutdown() {

	defer srv.wait.Done()
	if err := srv.e.Shutdown(context.Background()); err != nil {
		log.WithError(err).Error("Failed to shutdown server.")
	}
}
