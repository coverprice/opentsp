// Copyright 2015 The Sporting Exchange Limited. All rights reserved.
// Use of this source code is governed by a free license that can be
// found in the LICENSE file.

// tsp-poller is like tsp-forwarder except it extracts remote data.
package main

import (
	"log"
	"os"

	_ "opentsp.org/internal/pprof"

	"opentsp.org/internal/collect"
	"opentsp.org/internal/flag"
	"opentsp.org/internal/logfile"
	"opentsp.org/internal/relay"
	"opentsp.org/internal/stats"
	"opentsp.org/internal/tsdb"
	"opentsp.org/internal/tsdb/filter"
)

var cfg *Config

func init() {
	flag.Parse(defaultConfigPath)
	cfg = load(flag.FilePath)
	w := logfile.Open(cfg.LogPath)
	log.SetOutput(w)
	if flag.DebugMode {
		collect.Debug = log.New(w, "debug: collect: ", 0)
		filter.Debug = log.New(w, "debug: filter: ", 0)
	}
	log.Print("start pid=", os.Getpid())
}

func main() {
	var (
		plugins = collect.NewPool(cfg.CollectPath)
		self    = stats.Self("tsp.poller.")
		joined  = tsdb.Join(plugins.C, self)
		final   = filter.Series(cfg.Filter, joined)
		relays  = relay.NewPool(cfg.Relay, final)
	)
	go Restart(func() {
		plugins.Kill()
	})
	relays.Broadcast()
}
