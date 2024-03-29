// This file will be generated to include all customer specific middlewares
package registry

import (
	"github.com/FoxComm/vulcand/plugin"
	"github.com/FoxComm/vulcand/plugin/cbreaker"
	"github.com/FoxComm/vulcand/plugin/connlimit"
	"github.com/FoxComm/vulcand/plugin/ratelimit"
	"github.com/FoxComm/vulcand/plugin/rewrite"
	"github.com/FoxComm/vulcand/plugin/trace"
)

func GetRegistry() *plugin.Registry {
	r := plugin.NewRegistry()

	specs := []*plugin.MiddlewareSpec{
		ratelimit.GetSpec(),
		connlimit.GetSpec(),
		rewrite.GetSpec(),
		cbreaker.GetSpec(),
		trace.GetSpec(),
	}

	for _, spec := range specs {
		if err := r.AddSpec(spec); err != nil {
			panic(err)
		}
	}

	return r
}
