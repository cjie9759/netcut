package config

import "flag"

func Init() {
	initFlag()
}

var Port string

func initFlag() {
	flag.StringVar(&Port, "p", "0.0.0.0:80", "web端口，默认0.0.0.0:80")
}

var (
	// TorSize      = 1024 * 10
	// TotleSize    = 1024 * 50
	// PushLimit    = 50
	// TorShortName = ""

	// GC_LIMIT = 100 * 1024 * 1024 * 1024 // 100g

	DbPath = "fa4.db"

// Group_id     = 938132468
// Group_vod_id = "563112441"
)
