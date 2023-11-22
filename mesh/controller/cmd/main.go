package main

import (
	"etrib5gc/logctx"
	"etrib5gc/mesh/controller"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Usage: "Load configuration from `FILE`",
	},
	cli.StringFlag{
		Name:  "log, l",
		Usage: "Output logs to `FILE`",
	},
}

var _srv *controller.Controller
var _logfile *os.File
var log logctx.LogWriter

func main() {
	fmt.Println("Service Controller starts running")

	app := cli.NewApp()
	app.Name = "controller"
	app.Usage = "B5gc service controller"
	app.Action = action
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		return
	}
	quit := make(chan struct{})
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigch
		fmt.Println("System interrupted")
		if _srv != nil {
			_srv.Terminate()
		}
		quit <- struct{}{}
	}()
	<-quit
	t := time.NewTimer(200 * time.Millisecond)
	<-t.C
	log.Info("BYE")
	//close log file
	if _logfile != nil {
		_logfile.Close()
	}
}

func action(c *cli.Context) (err error) {
	//open log file
	logfilename := c.String("log")
	if len(logfilename) > 0 {
		if _logfile, err = os.OpenFile(logfilename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644); err != nil {
			fmt.Println("Failed to create logfile " + logfilename)
			return
		}
	}

	//read config
	var cfg controller.Config
	filename := c.String("config")
	if cfg, err = controller.LoadConfig(filename); err != nil {
		log.Errorf("Fail to parse controller configuration: %s", err.Error())
		return
	}

	//init logger
	loglevel := logctx.DEFAULT_LOG_LEVEL
	if cfg.LogLevel != nil {
		loglevel = *cfg.LogLevel
	}

	logctx.Init(_logfile, logctx.LogLevel(loglevel), logctx.Fields{
		"nf": "mesh-ctrl",
	})

	log = logctx.WithFields(logctx.Fields{
		"mod": "cmd",
	})

	//create the controller server
	if _srv, err = controller.New(&cfg); err != nil {
		log.Errorf("Fail to create the controller service: %s", err.Error())
		return
	}

	if err = _srv.Start(); err != nil {
		log.Errorf("Fail to start controller service: %s", err.Error())
		return
	}

	return
}
