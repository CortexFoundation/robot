package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CortexFoundation/CortexTheseus/log"
	"github.com/CortexFoundation/torrentfs/params"

	"github.com/CortexFoundation/robot"
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))

	cfg := &params.DefaultConfig
	cfg.DataDir = ".storage"
	cfg.RpcURI = "http://127.0.0.1:8545"

	var mm robot.IMonitor
	if m, err := robot.New(cfg, true, false, false, nil); err != nil {
		panic(err)
	} else {
		mm = m
		if err := mm.Start(); err != nil {
			log.Error("start failed", "err", err)
			panic(err)
		}
		defer mm.Stop()

		//m.SwitchService(robot.SRV_RECORD)
		go func() {
			for {
				//mm.SwitchService(robot.SRV_RECORD)
				//time.Sleep(30 * time.Second)
				//mm.SwitchService(robot.SRV_MODEL)
				time.Sleep(30 * time.Second)
			}
		}()

		var c = make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
	}
}
