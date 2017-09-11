// Copyright 2017 Mark Nevill. All Rights Reserved.
// See LICENSE for licensing terms.

package http_prometheus

import (
	"net/http"

	"github.com/mwitkow/go-httpwares/tags"
)

type meta struct {
	name, handler, method, host, path string
}

func reqMeta(req *http.Request, opts *options, inbound bool) *meta {
	m := &meta{name: opts.name}
	if m.name == "" && inbound {
		v := http_ctxtags.ExtractInbound(req).Values()[http_ctxtags.TagForCallService]
		m.name, _ = v.(string)
	}
	if m.name == "" && !inbound {
		v := http_ctxtags.ExtractOutbound(req).Values()[http_ctxtags.TagForCallService]
		m.name, _ = v.(string)
	}
	if inbound {
		v := http_ctxtags.ExtractInbound(req).Values()[http_ctxtags.TagForHandlerName]
		hname, _ := v.(string)
		if hname != "" {
			v := http_ctxtags.ExtractInbound(req).Values()[http_ctxtags.TagForHandlerGroup]
			hgroup, _ := v.(string)
			if hgroup == "" {
				hgroup = "unknown"
			}
			m.handler = hgroup + "." + hname
		}
	}
	if opts.hosts {
		m.host = req.URL.Host
		if m.host == "" {
			m.host = req.Host
		}
	}
	if opts.paths {
		m.path = req.URL.Path
	}
	return m
}