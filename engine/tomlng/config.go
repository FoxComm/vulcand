package tomlng

import (
	"github.com/FoxComm/vulcand/engine"
)

type EngineTomlConfig struct {
	Listeners   map[string]engine.Listener
	Frontends   map[string]Frontend
	Backends    map[string]Backend
	Servers     map[string][]engine.Server
	Middlewares map[string]engine.Middleware
}

type Frontend struct {
	engine.Frontend
	Settings engine.HTTPFrontendSettings
}

type Backend struct {
	engine.Backend
	Settings engine.HTTPBackendSettings
}

// type Server struct {
// 	Id  string
// 	URL string
// }

// type Middleware struct {
// 	Id string
// }
