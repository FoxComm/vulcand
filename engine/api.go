package engine

import (
	"github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/mailgun/scroll"
)

type EngineApiHandlers interface {
	AddApiHandlers(app *scroll.App)
}
