package tomlng

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/FoxComm/vulcand/engine"
	"github.com/FoxComm/vulcand/engine/test"
	"github.com/FoxComm/vulcand/log"
	"github.com/FoxComm/vulcand/plugin/registry"

	. "github.com/FoxComm/vulcand/Godeps/_workspace/src/gopkg.in/check.v1"
)

func TestToml(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&TomlSuite{})

type TomlSuite struct {
	suite   test.EngineSuite
	confDir string
	stopC   chan bool
}

func (s *TomlSuite) SetUpSuite(c *C) {
	log.EnsureLoggerExist("console", "INFO")
}

func (s *TomlSuite) SetUpTest(c *C) {
	var err error
	s.confDir, err = ioutil.TempDir(os.TempDir(), "fc_tomltest")
	c.Assert(err, IsNil)
	fmt.Printf("Config dir %s created %+v\n", s.confDir, err)

	engine, err := New(registry.GetRegistry(),
		Options{
			MainConfigFilepath: "",
			ConfigPaths:        []string{s.confDir},
			WatchConfigChanges: true,
		})
	if err != nil {
		c.Fatalf("Error while creating toml engine: %s", err.Error())
		return
	}

	tomlEngine, ok := engine.(*TomlNg)
	c.Assert(ok, Equals, true)
	tomlEngine.ReadConfig(strings.NewReader(test.TomlCfgDefaultListener))
	tomlEngine.ReadConfig(strings.NewReader(test.TomlCfgOriginFrontend))
	tomlEngine.ReadConfig(strings.NewReader(test.TomlCfgOriginBackend))

	s.suite.ChangesC = make(chan interface{})
	s.stopC = make(chan bool)
	go engine.Subscribe(s.suite.ChangesC, s.stopC)
	s.suite.Engine = engine
}

func (s *TomlSuite) TearDownTest(c *C) {
	close(s.stopC)
	s.suite.Engine.Close()
	// err := os.RemoveAll(s.confDir)
	// c.Assert(err, IsNil)
}

func (s *TomlSuite) newConfigFile(name string) (*os.File, error) {
	path := strings.Join([]string{s.confDir, name + ".toml"}, string(os.PathSeparator))
	return os.Create(path)

}

func (s *TomlSuite) TestHostCRUD(c *C) {
	s.suite.HostCRUD(c)
}

func (s *TomlSuite) TestHostWithKeyPair(c *C) {
	s.suite.HostWithKeyPair(c)
}

func (s *TomlSuite) TestHostUpsertKeyPair(c *C) {
	s.suite.HostUpsertKeyPair(c)
}

func (s *TomlSuite) TestHostWithOCSP(c *C) {
	s.suite.HostWithOCSP(c)
}

func (s *TomlSuite) TestListenerCRUD(c *C) {
	s.suite.ListenerCRUD(c)
}

func (s *TomlSuite) TestListenerSettingsCRUD(c *C) {
	s.suite.ListenerSettingsCRUD(c)
}

func (s *TomlSuite) TestBackendCRUD(c *C) {
	s.suite.BackendCRUD(c)
}

func (s *TomlSuite) TestBackendDeleteUsed(c *C) {
	s.suite.BackendDeleteUsed(c)
}

func (s *TomlSuite) TestServerCRUD(c *C) {
	s.suite.ServerCRUD(c)
}

func (s *TomlSuite) TestFrontendCRUD(c *C) {
	s.suite.FrontendCRUD(c)
}

func (s *TomlSuite) TestFrontendBadBackend(c *C) {
	s.suite.FrontendBadBackend(c)
}

func (s *TomlSuite) TestMiddlewareCRUD(c *C) {
	s.suite.MiddlewareCRUD(c)
}

func (s *TomlSuite) TestMiddlewareBadFrontend(c *C) {
	s.suite.MiddlewareBadFrontend(c)
}

func (s *TomlSuite) TestMiddlewareBadType(c *C) {
	s.suite.MiddlewareBadType(c)
}

func (s *TomlSuite) TestNewAndModifyFile(c *C) {
	log.Infof("f started")
	f, err := s.newConfigFile("server")
	defer f.Close()
	log.Infof("config file: %s", f.Name())
	c.Assert(err, IsNil)
	bk := engine.BackendKey{Id: "origin_frontend"}

	servers, err := s.suite.Engine.GetServers(bk)
	c.Assert(err, IsNil)
	c.Assert(len(servers), Equals, 0)

	time.Sleep(100 * time.Millisecond)
	f.WriteString(`[Servers]
    [[Servers.origin_frontend]]
    URL = "http://localhost:8080"
    `)
	f.Sync()
	// wait for reload event
	c.Assert(waitChannel(s.suite.ChangesC, 100*time.Millisecond), IsNil)

	servers, err = s.suite.Engine.GetServers(bk)
	c.Assert(err, IsNil)
	c.Assert(len(servers), Equals, 1)

	f.WriteString(`
	   [[Servers.origin_frontend]]
	   URL = "http://localhost:8085"
	   `)
	f.Sync()
	c.Assert(waitChannel(s.suite.ChangesC, 100*time.Millisecond), IsNil)

	servers, err = s.suite.Engine.GetServers(bk)
	c.Assert(err, IsNil)
	c.Assert(len(servers), Equals, 2)

	f.Close()
	os.Remove(f.Name())

	c.Assert(waitChannel(s.suite.ChangesC, 100*time.Millisecond), IsNil)
	c.Assert(waitChannel(s.suite.ChangesC, 100*time.Millisecond), IsNil)

	servers, err = s.suite.Engine.GetServers(bk)
	c.Assert(err, IsNil)
	c.Assert(len(servers), Equals, 0)
}

func waitChannel(ch chan interface{}, d time.Duration) error {
	select {
	case <-ch:
		return nil
	case <-time.After(d):
		return fmt.Errorf("Timeout on channel: %+v", ch)
	}
}
