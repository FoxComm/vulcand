package tomlng

import (
	"github.com/BurntSushi/toml"
	"github.com/FoxComm/vulcand/engine"
	//	"github.com/FoxComm/vulcand/plugin"
)

type EngineTomlConfig struct {
	Listeners   map[string]engine.Listener
	Frontends   map[string]Frontend
	Backends    map[string]Backend
	Servers     map[string][]engine.Server
	Middlewares map[string]MiddlewareFrontend
}

type Frontend struct {
	engine.Frontend
	Settings    engine.HTTPFrontendSettings
	Middlewares []FrontendMiddleware
}

type FrontendMiddleware struct {
	MiddlewareId string
	Priority     int
}

type Backend struct {
	engine.Backend
	Settings engine.HTTPBackendSettings
}

// type Server struct {
// 	Id  string
// 	URL string
// }

type MiddlewareFrontend struct {
	Type       string
	Middleware toml.Primitive
}
