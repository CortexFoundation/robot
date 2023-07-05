package main

import (
	"github.com/CortexFoundation/robot"
	"github.com/CortexFoundation/torrentfs/params"
)

func main() {
	cfg := &params.DefaultConfig
	cfg.DataDir = "test"
	if m, err := robot.New(cfg, true, false, false, nil); err != nil {
		panic(err)
	} else {
		m.Start()
		m.Stop()
	}
}
