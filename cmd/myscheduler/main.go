package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danieeelfr/myscheduler/internal/config"
	controller "github.com/danieeelfr/myscheduler/internal/controller/server"
	"github.com/danieeelfr/myscheduler/pkg/wait"
	"github.com/sirupsen/logrus"
)

var (
	log            = logrus.WithField("package", "main.myscheduler")
	wg             = wait.New()
	waitToShutdown time.Duration
)

func init() {
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}

	logrus.SetLevel(level)
}

func main() {
	log.Info("starting...")

	cfg := config.New(config.MySchedulerApp)

	waitToShutdown = time.Duration(cfg.GetMySchedulerConf().WaitToShutdown) * time.Second

	ctrlHTTP, err := controller.New(cfg, wg)
	if err != nil {
		log.WithError(err).Fatal("couldn't run the application")
	}

	wg.Add()

	shutdownSignal(ctrlHTTP)

	if err := ctrlHTTP.Start(); err != nil {
		log.WithError(err).Fatal("fail starting http server")
	}

	wg.Wait()

	log.Infof("Finishing %s...", config.MySchedulerApp)

}

func shutdownSignal(ctrlHTTP controller.Interactor) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-signalChannel
		switch sig {
		case syscall.SIGTERM, syscall.SIGINT:
			log.Infof("Interruption request. Signal: [%v].", sig)
			wg.Block()
			ctrlHTTP.Shutdown()

			log.Infof("Waiting [%v] for open processes.", waitToShutdown)
			time.Sleep(waitToShutdown)

			log.Infof("Finishing...")
			for wg.Done() {
				log.Infof("Ignoring open process to kill...")
			}
		}
	}()
}
