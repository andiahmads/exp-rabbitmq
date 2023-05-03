package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Numworkers int    `envconfig:"num_workers" default:"2"`
	QueueName  string `envconfig:"queue_name" default:"hello"`
}

func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg

}
