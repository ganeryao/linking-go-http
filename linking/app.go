package linking

import (
	"github.com/alecthomas/log4go"
)

// App is the base linking struct
type App struct {
	configured  bool
	debug       bool
}

var (
	app = &App{
		debug:       false,
		configured:  false,
	}
)

// Configure configures the linking
func Configure(
	debug bool,
) {
	if app.configured {
		_ = log4go.Warn("lk http configured twice!")
	}
	app.configured = true
	app.debug = debug
}

