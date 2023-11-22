package main

import (
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	"etrib5gc/nfs/amf/config"
	"etrib5gc/nfs/amf/service"
	"fmt"
	"os"
	"os/signal"
	"sync"
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

var _nf *service.AMF
var _logfile *os.File
var log logctx.LogWriter

func main() {
	fmt.Println("B5gc AMF starts running")

	app := cli.NewApp()
	app.Name = "amf"
	app.Usage = "5G Access and Mobility Management Function (AMF)"
	app.Action = action
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		return
	}
	var wg sync.WaitGroup
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		defer mesh.Terminate()
		defer wg.Done()
		<-sigch
		fmt.Println("System interrupted")
		if _nf != nil {
			_nf.Terminate()
		}
	}()
	wg.Wait()

	t := time.NewTimer(200 * time.Millisecond)
	<-t.C

	fmt.Println("BYE")

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
	var cfg config.AmfConfig
	configfile := c.String("config")
	if cfg, err = config.LoadConfig(configfile); err != nil {
		fmt.Println("Fail to parse AMF configuration: " + err.Error())
		return
	}

	//init logger
	loglevel := logctx.DEFAULT_LOG_LEVEL
	if cfg.LogLevel != nil {
		loglevel = *cfg.LogLevel
	}

	logctx.Init(_logfile, loglevel, logctx.Fields{
		"nf": "amf",
	})

	log = logctx.WithFields(logctx.Fields{
		"mod": "cmd",
	})

	//create the AMF
	if _nf, err = service.New(&cfg); err != nil {
		log.Errorf("Fail to create AMF service: %s", err.Error())
		return
	}

	if err = mesh.Init(&cfg.Mesh, _nf); err != nil {
		log.Errorf("Fail to init mesh agent: %s", err.Error())
		return
	}

	if err = _nf.Start(); err != nil {
		log.Errorf("Fail to start AMF service: %s", err.Error())
		return
	}

	return
}
