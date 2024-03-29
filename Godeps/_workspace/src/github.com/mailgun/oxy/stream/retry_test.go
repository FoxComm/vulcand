package stream

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/mailgun/oxy/forward"
	"github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/mailgun/oxy/roundrobin"
	"github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/mailgun/oxy/testutils"
	"github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/mailgun/oxy/utils"

	. "github.com/FoxComm/vulcand/Godeps/_workspace/src/gopkg.in/check.v1"
)

type RTSuite struct{}

var _ = Suite(&RTSuite{})

func (s *RTSuite) TestSuccess(c *C) {
	srv := testutils.NewHandler(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})
	defer srv.Close()

	lb, rt := new(c, `IsNetworkError() && Attempts() <= 2`)

	proxy := httptest.NewServer(rt)
	defer proxy.Close()

	lb.UpsertServer(testutils.ParseURI(srv.URL))

	re, body, err := testutils.Get(proxy.URL)
	c.Assert(err, IsNil)
	c.Assert(re.StatusCode, Equals, http.StatusOK)
	c.Assert(string(body), Equals, "hello")
}

func (s *RTSuite) TestRetryOnError(c *C) {
	srv := testutils.NewHandler(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})
	defer srv.Close()

	lb, rt := new(c, `IsNetworkError() && Attempts() <= 2`)

	proxy := httptest.NewServer(rt)
	defer proxy.Close()

	lb.UpsertServer(testutils.ParseURI("http://localhost:64321"))
	lb.UpsertServer(testutils.ParseURI(srv.URL))

	re, body, err := testutils.Get(proxy.URL, testutils.Body("some request parameters"))
	c.Assert(err, IsNil)
	c.Assert(re.StatusCode, Equals, http.StatusOK)
	c.Assert(string(body), Equals, "hello")
}

func (s *RTSuite) TestRetryExceedAttempts(c *C) {
	srv := testutils.NewHandler(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})
	defer srv.Close()

	lb, rt := new(c, `IsNetworkError() && Attempts() <= 2`)

	proxy := httptest.NewServer(rt)
	defer proxy.Close()

	lb.UpsertServer(testutils.ParseURI("http://localhost:64321"))
	lb.UpsertServer(testutils.ParseURI("http://localhost:64322"))
	lb.UpsertServer(testutils.ParseURI("http://localhost:64323"))
	lb.UpsertServer(testutils.ParseURI(srv.URL))

	re, _, err := testutils.Get(proxy.URL)
	c.Assert(err, IsNil)
	c.Assert(re.StatusCode, Equals, http.StatusBadGateway)
}

func new(c *C, p string) (*roundrobin.RoundRobin, *Streamer) {
	logger := utils.NewFileLogger(os.Stdout, utils.INFO)
	// forwarder will proxy the request to whatever destination
	fwd, err := forward.New(forward.Logger(logger))
	c.Assert(err, IsNil)

	// load balancer will round robin request
	lb, err := roundrobin.New(fwd)
	c.Assert(err, IsNil)

	// stream handler will forward requests to redirect, make sure it uses files
	st, err := New(lb, Logger(logger), Retry(p), MemRequestBodyBytes(1))
	c.Assert(err, IsNil)

	return lb, st
}
