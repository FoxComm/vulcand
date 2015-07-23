// package TomlNgng provides in toml engine implementation
package tomlng

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/FoxComm/vulcand/engine"
	"github.com/FoxComm/vulcand/plugin"

	"github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/BurntSushi/toml"
	"github.com/FoxComm/vulcand/Godeps/_workspace/src/gopkg.in/fsnotify.v1"

	"github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/mailgun/scroll"
	"github.com/FoxComm/vulcand/log"

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

	options        Options
	tomlConfig     EngineTomlConfig
	tomlMeta       toml.MetaData
	tomlSyncerLock sync.Mutex

	tomlWatcher *fsnotify.Watcher
}

type Options struct {
	MainConfigFilepath string
	ConfigPaths        []string
	WatchConfigChanges bool
}

func New(r *plugin.Registry, options Options) (engine.Engine, error) {
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
		options:          options,
	}
	if options.WatchConfigChanges {
		ng.watchConfigFiles()
	}

	for _, p := range options.ConfigPaths {
		if err := ng.AddConfigPath(p); err != nil {
			return nil, err
		}
	}

	err := ng.reloadConfig()
	return ng, err
}

func (m *TomlNg) emit(val interface{}) {
	select {
	case m.ChangesC <- val:
	default:
	}
}

func (m *TomlNg) Close() {
	if m.options.WatchConfigChanges {
		m.tomlWatcher.Close()
	}
}

func (m *TomlNg) AddConfigPath(in string) error {
	if in == "" {
		return nil
	}
	absin, err := pathAbs(in)
	if err != nil {
		return err
	}

	configPathExists, err := pathExists(absin)
	if err != nil {
		return err
	}
	if !configPathExists {
		return fmt.Errorf("Config path: %s not exists", absin)
	}

	if !stringInSlice(absin, m.options.ConfigPaths) && configPathExists {
		if m.options.WatchConfigChanges {
			m.tomlWatcher.Add(absin)
		}
		m.options.ConfigPaths = append(m.options.ConfigPaths, absin)
	}

	return nil
}

func (m *TomlNg) watchConfigFiles() (err error) {
	m.tomlWatcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func() {
		opsWatched := fsnotify.Create | fsnotify.Write | fsnotify.Remove
		for {
			select {
			case event := <-m.tomlWatcher.Events:
				if event.Op&opsWatched > 0 && path.Ext(event.Name) == ".toml" {
					err = m.reloadConfig()
					if err != nil {
						log.Errorf("Error while decoding new config file: %s, %s", event.Name, err.Error())
						continue
					}

					continue
				}
			case err := <-m.tomlWatcher.Errors:
				log.Errorf("error: %v", err)
			}
		}
	}()
	return
}

func (m *TomlNg) loadConfig(config *EngineTomlConfig) error {
	var err error
	if m.options.MainConfigFilepath != "" {
		if m.tomlMeta, err = toml.DecodeFile(m.options.MainConfigFilepath, config); err != nil {
			return err
		}
	}

	for _, configpath := range m.options.ConfigPaths {
		configFiles, err := filepath.Glob(path.Join(configpath, "*.toml"))
		if err != nil {
			return err
		}
		for _, cfg := range configFiles {
			log.Infof("Load toml config from %s", cfg)
			_, err = toml.DecodeFile(cfg, config)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//  reloadConfig do 3 steps
// 1. First memoize current config
// 2. decode new values into new tomlConfig structure and replace old one.
// 3. Add new entities to state and delete obosoletes using info from steps 1,2
func (m *TomlNg) reloadConfig() error {
	m.tomlSyncerLock.Lock()
	defer m.tomlSyncerLock.Unlock()
	// Memoize current config
	existingListeners := mapStringKeys(m.tomlConfig.Listeners)
	existingBackends := mapStringKeys(m.tomlConfig.Backends)
	existingMiddlewares := mapStringKeys(m.tomlConfig.Middlewares)

	existingServers := []engine.ServerKey{}
	for key, servers := range m.tomlConfig.Servers {
		for _, srv := range servers {
			sk := engine.ServerKey{Id: serverKey(key, srv), BackendKey: engine.BackendKey{Id: key}}
			existingServers = append(existingServers, sk)
		}
	}
	// name => list of middlewares
	existingFrontends := map[string][]string{}
	for key, f := range m.tomlConfig.Frontends {
		existingFrontends[key] = []string{}
		for _, m := range f.Middlewares {
			existingFrontends[key] = append(existingFrontends[key], m.MiddlewareId)
		}
	}

	// decode config into new structure
	// we can't use there m.tomlConfig
	// because loadConfig persist old values
	var newConfig EngineTomlConfig
	if err := m.loadConfig(&newConfig); err != nil {
		return err
	}
	m.tomlConfig = newConfig

	// sync state
	if err := m.syncListeners(existingListeners); err != nil {
		return err
	}

	if err := m.syncMiddlewares(existingMiddlewares); err != nil {
		return err
	}

	if err := m.syncBackends(existingBackends); err != nil {
		return err
	}

	if err := m.syncFrontends(existingFrontends); err != nil {
		return err
	}

	if err := m.syncServers(existingServers); err != nil {
		return err
	}

	return nil
}

func (m *TomlNg) syncListeners(keysCurrent []string) error {
	// First, add listeners, that should be added
	for id, lr := range m.tomlConfig.Listeners {
		// Create listener instance from config
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

		if err := m.UpsertListener(*newLr); err != nil {
			return err
		}
	}

	// Second, remove listeners that should not be there any more
	for _, key := range keysCurrent {
		if _, ok := m.tomlConfig.Listeners[key]; !ok {
			if err := m.DeleteListener(engine.ListenerKey{Id: key}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *TomlNg) syncFrontends(currentCfg map[string][]string) error {
	for id, rf := range m.tomlConfig.Frontends {
		rf.Id = id
		if rf.Type != engine.HTTP {
			return fmt.Errorf("Unsupported frontend type: %v", rf.Type)
		}

		f, err := engine.NewHTTPFrontend(rf.Id, rf.BackendId, rf.Route, rf.Settings)
		if err != nil {
			return err
		}
		if err := m.UpsertFrontend(*f, 0); err != nil {
			return err
		}

		newMiddlewareList := []string{}
		for _, mRef := range rf.Middlewares {
			middleware, ok := m.KnownMiddlewares[mRef.MiddlewareId]
			if !ok {
				return fmt.Errorf("Middleware %s not loaded or not exist", mRef.MiddlewareId)
			}
			copyOfMiddleware := middleware
			copyOfMiddleware.Priority = mRef.Priority

			if err := m.UpsertMiddleware(engine.FrontendKey{Id: rf.Id}, copyOfMiddleware, 0); err != nil {
				return err
			}
			newMiddlewareList = append(newMiddlewareList, mRef.MiddlewareId)
		}

		if currentMiddlewareList, ok := currentCfg[rf.Id]; ok {
			for _, key := range currentMiddlewareList {
				if !stringInSlice(key, newMiddlewareList) {
					if err := m.DeleteMiddleware(engine.MiddlewareKey{FrontendKey: f.GetKey(), Id: key}); err != nil {
						return err
					}
				}
			}
		}

	}

	// Second, remove frontends that should not be there any more
	keysCurrent := []string{}
	for k, _ := range currentCfg {
		keysCurrent = append(keysCurrent, k)
	}

	for _, key := range keysCurrent {
		if _, ok := m.tomlConfig.Frontends[key]; !ok {
			if err := m.DeleteFrontend(engine.FrontendKey{Id: key}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *TomlNg) syncBackends(keysCurrent []string) error {
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
		if err := m.UpsertBackend(*b); err != nil {
			return err
		}
	}
	// Second, remove backends that should not be there any more
	for _, key := range keysCurrent {
		if _, ok := m.tomlConfig.Backends[key]; !ok {
			if err := m.DeleteBackend(engine.BackendKey{Id: key}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *TomlNg) syncServers(keysCurrent []engine.ServerKey) error {
	for id, servers := range m.tomlConfig.Servers {
		bkey := engine.BackendKey{Id: id}
		for _, s := range servers {
			skey := serverKey(id, s)
			server, err := engine.NewServer(skey, s.URL)
			if err != nil {
				return err
			}
			if err := m.UpsertServer(bkey, *server, 0); err != nil {
				return err
			}
		}
	}
	// Second, remove servers that should not be there any more
	newKeys := []string{}
	for bkey, servers := range m.tomlConfig.Servers {
		for _, srv := range servers {
			newKey := serverKey(bkey, srv)
			newKeys = append(newKeys, newKey)
		}
	}

	for _, key := range keysCurrent {
		if !stringInSlice(key.Id, newKeys) {
			if err := m.DeleteServer(key); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *TomlNg) syncMiddlewares(keysCurrent []string) error {
	for id, ms := range t.tomlConfig.Middlewares {
		spec := t.Registry.GetSpec(ms.Type)
		if spec == nil {
			return fmt.Errorf("middleware of type %s is not supported", ms.Type)
		}
		m, err := spec.FromToml(ms.Middleware, &t.tomlMeta)

		if f, ok := m.(engine.InitMiddlewareEngine); ok {
			f.InitEngine(t)
		}

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

	// Second, remove middlewares that should not be there any more
	for _, key := range keysCurrent {
		if _, ok := t.tomlConfig.Middlewares[key]; !ok {
			delete(t.KnownMiddlewares, key)
		}
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
	key := engine.ListenerKey{Id: l.Id}

	defer func() {
		m.emit(&engine.ListenerUpserted{Listener: l})
	}()

	m.Listeners[key] = l
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
		return &engine.NotFoundError{Message: fmt.Sprintf("Can't upsert middleware %s, Frontend %v' not found", md.Id, fk)}
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

func (m *TomlNg) AddApiHandlers(app *scroll.App) {

	app.AddHandler(scroll.Spec{
		Paths:   []string{"/v2/engine/reloadconfig"},
		Methods: []string{"POST"},
		Handler: m.apiReloadConfig,
	})
}

func (m *TomlNg) apiReloadConfig(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	err := m.reloadConfig()
	if err != nil {
		return nil, err
	}
	return "ok", err
}

func serverKey(backendId string, server engine.Server) string {
	urlPart := strings.Replace(server.URL, "://", "_", -1)
	urlPart = strings.Replace(urlPart, ":", "_", -1)
	return fmt.Sprintf("%s_%s", backendId, urlPart)
}
