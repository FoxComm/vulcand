// package TomlNgng provides in toml engine implementation
package tomlng

import (
	"fmt"

	"time"

	"github.com/FoxComm/vulcand/engine"
	"github.com/FoxComm/vulcand/plugin"

	"github.com/BurntSushi/toml"
	"github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/mailgun/log"

	// "io/ioutil"
	// "os"
	// "github.com/naoina/toml"
)

type TomlNg struct {
	Hosts     map[engine.HostKey]engine.Host
	Frontends map[engine.FrontendKey]engine.Frontend
	Backends  map[engine.BackendKey]engine.Backend
	Listeners map[engine.ListenerKey]engine.Listener

	Middlewares      map[engine.FrontendKey][]engine.Middleware
	KnownMiddlewares map[string]engine.Middleware
	Servers          map[engine.BackendKey][]engine.Server

	Registry *plugin.Registry
	ChangesC chan interface{}
	ErrorsC  chan error

	tomlConfigPath string
	tomlConfig     EngineTomlConfig
	tomlMeta       toml.MetaData
}

func New(configPath string, r *plugin.Registry) (engine.Engine, error) {
	ng := &TomlNg{
		Hosts:     map[engine.HostKey]engine.Host{},
		Frontends: map[engine.FrontendKey]engine.Frontend{},
		Backends:  map[engine.BackendKey]engine.Backend{},

		Listeners:        map[engine.ListenerKey]engine.Listener{},
		Middlewares:      map[engine.FrontendKey][]engine.Middleware{},
		KnownMiddlewares: map[string]engine.Middleware{},
		Servers:          map[engine.BackendKey][]engine.Server{},
		Registry:         r,
		ChangesC:         make(chan interface{}, 1000),
		ErrorsC:          make(chan error),
		tomlConfigPath:   configPath,
	}
	err := ng.LoadConfig()
	return ng, err
}

func (m *TomlNg) emit(val interface{}) {
	select {
	case m.ChangesC <- val:
	default:
	}
}

func (m *TomlNg) Close() {
}

func (m *TomlNg) LoadConfig() error {
	var err error
	if m.tomlMeta, err = toml.DecodeFile(m.tomlConfigPath, &m.tomlConfig); err != nil {
		return err
	}

	fmt.Printf("TOML: %+v\n", m.tomlConfig)

	if err := m.loadListeners(); err != nil {
		return err
	}

	if err := m.loadMiddlewares(); err != nil {
		return err
	}

	if err := m.loadFrontends(); err != nil {
		return err
	}

	if err := m.loadBackends(); err != nil {
		return err
	}

	if err := m.loadServers(); err != nil {
		return err
	}

	return nil

}

func (m *TomlNg) loadListeners() error {
	for id, lr := range m.tomlConfig.Listeners {
		lr.Id = id
		if lr.Protocol == engine.HTTPS && lr.Settings != nil {
			if _, err := engine.NewTLSConfig(&lr.Settings.TLS); err != nil {
				return err
			}
		}

		newLr, err := engine.NewListener(lr.Id, lr.Protocol, lr.Address.Network, lr.Address.Address, lr.Scope, lr.Settings)
		if err != nil {
			return err
		}
		key := engine.ListenerKey{Id: id}
		m.Listeners[key] = *newLr
	}
	return nil
}

func (m *TomlNg) loadFrontends() error {
	for id, rf := range m.tomlConfig.Frontends {
		rf.Id = id
		if rf.Type != engine.HTTP {
			return fmt.Errorf("Unsupported frontend type: %v", rf.Type)
		}

		f, err := engine.NewHTTPFrontend(rf.Id, rf.BackendId, rf.Route, rf.Settings)
		if err != nil {
			return err
		}
		key := engine.FrontendKey{Id: id}
		m.Frontends[key] = *f
		for _, mRef := range rf.Middlewares {
			middleware, ok := m.KnownMiddlewares[mRef.MiddlewareId]
			copyOfMiddleware := middleware
			copyOfMiddleware.Priority = mRef.Priority
			if !ok {
				return fmt.Errorf("Middleware %s not loaded or not exist", mRef.MiddlewareId)
			}
			m.Middlewares[key] = append(m.Middlewares[key], copyOfMiddleware)
		}
	}
	return nil
}

func (m *TomlNg) loadBackends() error {
	for id, rb := range m.tomlConfig.Backends {
		rb.Id = id

		if rb.Type != engine.HTTP {
			return fmt.Errorf("Unsupported backend type %v", rb.Type)
		}

		if rb.Settings.TLS != nil {
			if _, err := engine.NewTLSConfig(rb.Settings.TLS); err != nil {
				return err
			}
		}

		b, err := engine.NewHTTPBackend(rb.Id, rb.Settings)
		if err != nil {
			return err
		}
		key := engine.BackendKey{Id: id}
		m.Backends[key] = *b
	}
	return nil
}

func (m *TomlNg) loadServers() error {
	for id, servers := range m.tomlConfig.Servers {
		bkey := engine.BackendKey{Id: id}
		for i, s := range servers {
			skey := fmt.Sprintf("%ssrv%d", id, i)
			server, err := engine.NewServer(skey, s.URL)
			if err != nil {
				return err
			}
			m.Servers[bkey] = append(m.Servers[bkey], *server)
		}
	}
	return nil
}

func (t *TomlNg) loadMiddlewares() error {
	for id, ms := range t.tomlConfig.Middlewares {
		spec := t.Registry.GetSpec(ms.Type)
		if spec == nil {
			return fmt.Errorf("middleware of type %s is not supported", ms.Type)
		}
		m, err := spec.FromToml(ms.Middleware, &t.tomlMeta)
		if err != nil {
			return err
		}
		middleware := engine.Middleware{
			Id:         id,
			Type:       ms.Type,
			Middleware: m,
			Priority:   0,
		}
		t.KnownMiddlewares[id] = middleware
	}
	return nil
}

func (m *TomlNg) GetRegistry() *plugin.Registry {
	return m.Registry
}

func (m *TomlNg) GetHosts() ([]engine.Host, error) {
	out := make([]engine.Host, 0, len(m.Hosts))
	for _, h := range m.Hosts {
		out = append(out, h)
	}
	return out, nil
}

func (m *TomlNg) GetHost(k engine.HostKey) (*engine.Host, error) {
	h, ok := m.Hosts[k]
	if !ok {
		return nil, &engine.NotFoundError{}
	}
	return &h, nil
}

func (m *TomlNg) UpsertHost(h engine.Host) error {
	m.Hosts[engine.HostKey{Name: h.Name}] = h
	m.emit(&engine.HostUpserted{Host: h})
	return nil
}

func (m *TomlNg) DeleteHost(k engine.HostKey) error {
	if _, ok := m.Hosts[k]; !ok {
		return &engine.NotFoundError{}
	}
	delete(m.Hosts, k)
	m.emit(&engine.HostDeleted{HostKey: k})
	return nil
}

func (m *TomlNg) GetListeners() ([]engine.Listener, error) {
	out := make([]engine.Listener, 0, len(m.Listeners))
	for _, l := range m.Listeners {
		out = append(out, l)
	}
	return out, nil
}

func (m *TomlNg) GetListener(lk engine.ListenerKey) (*engine.Listener, error) {
	val, ok := m.Listeners[lk]
	if !ok {
		return nil, &engine.NotFoundError{}
	}
	return &val, nil
}

func (m *TomlNg) UpsertListener(l engine.Listener) error {
	defer func() {
		m.emit(&engine.ListenerUpserted{Listener: l})
	}()
	lk := engine.ListenerKey{l.Id}
	m.Listeners[lk] = l
	return nil
}

func (m *TomlNg) DeleteListener(lk engine.ListenerKey) error {
	if _, ok := m.Listeners[lk]; !ok {
		return &engine.NotFoundError{}
	}
	delete(m.Listeners, lk)
	m.emit(&engine.ListenerDeleted{ListenerKey: lk})
	return nil
}

func (m *TomlNg) GetFrontends() ([]engine.Frontend, error) {
	out := make([]engine.Frontend, 0, len(m.Frontends))
	for _, h := range m.Frontends {
		out = append(out, h)
	}
	return out, nil
}

func (m *TomlNg) GetFrontend(k engine.FrontendKey) (*engine.Frontend, error) {
	f, ok := m.Frontends[k]
	if !ok {
		return nil, &engine.NotFoundError{}
	}
	return &f, nil
}

func (m *TomlNg) UpsertFrontend(f engine.Frontend, d time.Duration) error {
	if _, ok := m.Backends[engine.BackendKey{Id: f.BackendId}]; !ok {
		return &engine.NotFoundError{Message: fmt.Sprintf("backend: %v not found", f.BackendId)}
	}
	m.Frontends[engine.FrontendKey{Id: f.Id}] = f
	m.emit(&engine.FrontendUpserted{Frontend: f})
	return nil
}

func (m *TomlNg) DeleteFrontend(fk engine.FrontendKey) error {
	if _, ok := m.Frontends[fk]; !ok {
		return &engine.NotFoundError{}
	}
	m.emit(&engine.FrontendDeleted{FrontendKey: fk})
	delete(m.Frontends, fk)
	return nil
}

func (m *TomlNg) GetMiddlewares(fk engine.FrontendKey) ([]engine.Middleware, error) {
	vals, ok := m.Middlewares[fk]
	if !ok {
		return []engine.Middleware{}, nil
	}
	return vals, nil
}

func (m *TomlNg) GetMiddleware(mk engine.MiddlewareKey) (*engine.Middleware, error) {
	vals, ok := m.Middlewares[mk.FrontendKey]
	if !ok {
		return nil, &engine.NotFoundError{Message: fmt.Sprintf("'%v' not found", mk.FrontendKey)}
	}
	for _, v := range vals {
		if v.Id == mk.Id {
			return &v, nil
		}
	}
	return nil, &engine.NotFoundError{Message: fmt.Sprintf("'%v' not found", mk)}
}

func (m *TomlNg) UpsertMiddleware(fk engine.FrontendKey, md engine.Middleware, d time.Duration) error {
	if _, ok := m.Frontends[fk]; !ok {
		return &engine.NotFoundError{Message: fmt.Sprintf("'%v' not found", fk)}
	}
	defer func() {
		m.emit(&engine.MiddlewareUpserted{FrontendKey: fk, Middleware: md})
	}()
	vals, ok := m.Middlewares[fk]
	if !ok {
		m.Middlewares[fk] = []engine.Middleware{md}
		return nil
	}
	for i, v := range vals {
		if v.Id == md.Id {
			vals[i] = md
			return nil
		}
	}
	vals = append(vals, md)
	m.Middlewares[fk] = vals
	return nil
}

func (m *TomlNg) DeleteMiddleware(mk engine.MiddlewareKey) error {
	vals, ok := m.Middlewares[mk.FrontendKey]
	if !ok {
		return &engine.NotFoundError{}
	}
	for i, v := range vals {
		if v.Id == mk.Id {
			vals = append(vals[:i], vals[i+1:]...)
			m.Middlewares[mk.FrontendKey] = vals
			m.emit(&engine.MiddlewareDeleted{MiddlewareKey: mk})
			return nil
		}
	}
	return &engine.NotFoundError{}
}

func (m *TomlNg) GetBackends() ([]engine.Backend, error) {
	out := make([]engine.Backend, 0, len(m.Backends))
	for _, h := range m.Backends {
		out = append(out, h)
	}
	return out, nil
}

func (m *TomlNg) GetBackend(bk engine.BackendKey) (*engine.Backend, error) {
	f, ok := m.Backends[bk]
	if !ok {
		return nil, &engine.NotFoundError{}
	}
	return &f, nil
}

func (m *TomlNg) UpsertBackend(b engine.Backend) error {
	m.emit(&engine.BackendUpserted{Backend: b})
	m.Backends[engine.BackendKey{Id: b.Id}] = b
	return nil
}

func (m *TomlNg) DeleteBackend(bk engine.BackendKey) error {
	for _, f := range m.Frontends {
		if f.BackendId == bk.Id {
			return fmt.Errorf("Backend is in use by %v", f)
		}
	}
	if _, ok := m.Backends[bk]; !ok {
		return &engine.NotFoundError{}
	}
	m.emit(&engine.BackendDeleted{BackendKey: bk})
	delete(m.Backends, bk)
	return nil
}

func (m *TomlNg) GetServers(bk engine.BackendKey) ([]engine.Server, error) {
	vals, ok := m.Servers[bk]
	if !ok {
		return []engine.Server{}, nil
	}
	return vals, nil
}

func (m *TomlNg) GetServer(sk engine.ServerKey) (*engine.Server, error) {
	vals, ok := m.Servers[sk.BackendKey]
	if !ok {
		return nil, &engine.NotFoundError{}
	}
	for _, v := range vals {
		if v.Id == sk.Id {
			return &v, nil
		}
	}
	return nil, &engine.NotFoundError{}
}

func (m *TomlNg) UpsertServer(bk engine.BackendKey, srv engine.Server, d time.Duration) error {
	defer func() {
		m.emit(&engine.ServerUpserted{BackendKey: bk, Server: srv})
	}()
	vals, ok := m.Servers[bk]
	if !ok {
		m.Servers[bk] = []engine.Server{srv}
		return nil
	}
	for i, v := range vals {
		if v.Id == srv.Id {
			m.Servers[bk][i] = srv
			return nil
		}
	}
	m.Servers[bk] = append(vals, srv)
	return nil
}

func (m *TomlNg) DeleteServer(sk engine.ServerKey) error {
	vals, ok := m.Servers[sk.BackendKey]
	if !ok {
		return &engine.NotFoundError{}
	}
	for i, v := range vals {
		if v.Id == sk.Id {
			vals = append(vals[:i], vals[i+1:]...)
			m.Servers[sk.BackendKey] = vals
			m.emit(&engine.ServerDeleted{ServerKey: sk})
			return nil
		}
	}
	return &engine.NotFoundError{}
}

func (m *TomlNg) Subscribe(changes chan interface{}, cancelC chan bool) error {
	for {
		select {
		case <-cancelC:
			return nil
		case change := <-m.ChangesC:
			log.Infof("Got change: %v", change)
			select {
			case changes <- change:
			case err := <-m.ErrorsC:
				log.Infof("Returning error: %v", err)
				return err
			}
		case err := <-m.ErrorsC:
			log.Infof("Returning error: %v", err)
			return err
		}
	}
}
