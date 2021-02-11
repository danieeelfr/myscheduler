package config

type Config struct {
	myschedulerConf *MySchedulerConf
}

type MySchedulerConf struct {
	HttpServerHost string
	WaitToShutdown uint
}

func (cfg *Config) GetMySchedulerConf() *MySchedulerConf { return cfg.myschedulerConf }
