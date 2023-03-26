package config

import "flag"

func Init() {
	initFlag()
}

var Port string

func initFlag() {
	flag.StringVar(&Port, "p", "0.0.0.0:10080", "web端口，默认0.0.0.0:10080")
	flag.Parse()
}

var (
	DbPath = ":memory:"
)
