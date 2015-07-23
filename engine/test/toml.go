package test

const (
	TomlCfgDefaultListener = `
[listeners]
  [listeners.default]
    protocol = "http"
  [listeners.default.address]
    network = "tcp"
    address = "0.0.0.0:18080"

`

	TomlCfgOriginFrontend = `
[frontends]
  [frontends.origin_frontend]
    Route = "PathRegexp("^/.*")"
    Type = "http"
    BackendId = "origin_frontend"
`
	TomlCfgOriginBackend = `
  [backends]

  [backends.origin_frontend]
    type = "http"
  [backends.origin_frontend.settings]
  [backends.origin_frontend.settings.tls]
    InsecureSkipVerify = true
    MinVersion = "VersionTLS10"
  [backends.origin_frontend.settings.keepalive]
    period = "60s"
`
)
