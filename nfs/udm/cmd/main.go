package main

import (
	"etrib5gc/logctx"
	"etrib5gc/mesh"
	"etrib5gc/nfs/udm/config"
	"etrib5gc/nfs/udm/service"
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

var _logfile *os.File
var log logctx.LogWriter
var _nf *service.UDM

func main() {
	fmt.Println("B5gc UDM starts runnings")

	app := cli.NewApp()
	app.Name = "udm"
	app.Usage = "Etri 5G Unified Datamanagement Function (UDM)"
	app.Action = action
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
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
	var cfg config.UdmConfig
	filename := c.String("config")
	if cfg, err = config.LoadConfig(filename); err != nil {
		fmt.Println("Fail to parse UDM configuration: " + err.Error())
		return
	}

	//init top level logger
	loglevel := logctx.DEFAULT_LOG_LEVEL
	if cfg.LogLevel != nil {
		loglevel = *cfg.LogLevel
	}
	logctx.Init(_logfile, loglevel, logctx.Fields{
		"nf": "udm",
	})
	log = logctx.WithFields(logctx.Fields{
		"mode": "cmd",
	})

	//create the UDM
	if _nf, err = service.New(&cfg); err != nil {
		log.Errorf("Fail to create UDM service: %s", err.Error())
		return
	}
	//create mesh agent
	if err = mesh.Init(&cfg.Mesh, _nf); err != nil {
		log.Errorf("Fail to init mesh agent: %s", err.Error())
		return
	}

	if err = _nf.Start(); err != nil {
		log.Errorf("Fail to start UDM service: %s", err.Error())
		return
	}

	return
}
