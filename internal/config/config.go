package config

import "github.com/sirupsen/logrus"

const (
	MySchedulerApp = "my-scheduler-app"
)

var (
	log = logrus.WithField("package", "config")
)

func New(app string) *Config {
	cfg := new(Config)

	switch app {
	case MySchedulerApp:
		cfg.myschedulerConf = &MySchedulerConf{
			HttpServerHost: "localhost:1313",
		}
	}

	return cfg
}
