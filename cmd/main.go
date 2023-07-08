package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CortexFoundation/CortexTheseus/log"
	"github.com/CortexFoundation/robot"
	"github.com/CortexFoundation/torrentfs/params"
)

func main() {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))
	glogger.Verbosity(log.LvlInfo)
	glogger.Vmodule("")
	log.Root().SetHandler(glogger)

	cfg := &params.DefaultConfig
	cfg.DataDir = ".storage"
	cfg.RpcURI = "http://127.0.0.1:8545"

	if m, err := robot.New(cfg, true, false, false, nil); err != nil {
		panic(err)
	} else {
		if err := m.Start(); err != nil {
			log.Error("start failed", "err", err)
			panic(err)
		}
		defer m.Stop()

		//m.SwitchService(robot.SRV_RECORD)
		go func() {
			for {
				m.SwitchService(robot.SRV_RECORD)
				time.Sleep(30 * time.Second)
				m.SwitchService(robot.SRV_MODEL)
				time.Sleep(30 * time.Second)
				//m.SwitchService(robot.SRV_RECORD)
			}
		}()

		var c = make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
	}
}
