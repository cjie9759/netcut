package config

import "flag"

func Init() {
	initFlag()
}

var Port string
var DbPath string

func initFlag() {
	flag.StringVar(&Port, "p", "0.0.0.0:10080", "web端口，默认0.0.0.0:10080")
	// flag.StringVar(&DbPath, "d", ":memory:", "dbpath，默认:memory:")
	flag.StringVar(&DbPath, "d", "./t.db", "dbpath，默认./t.db")
	flag.Parse()
}
