package mesh

import (
	"etrib5gc/common"
	"etrib5gc/mesh/httpdp"
	"etrib5gc/mesh/models"
	"etrib5gc/mesh/registry"
	"etrib5gc/sbi"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

func Init(config *MeshConfig, app App) (err error) {
	_initLog()
	if _mesh, err = newMesh(config, app); err != nil {
		_mesh = nil
	} else {
		rand.Seed(time.Now().UnixNano())
	}
	log.Info("mesh is initialized")
	return
}

func Terminate() (err error) {
	if _mesh != nil {
		err = _mesh.stop()
	}
	return
}

func Consumer(id models.ServiceName, match models.RouteMatch, stateless bool) (cli sbi.ConsumerClient, err error) {
	var matched registry.MatchedGroup
	if matched, err = _mesh.regcli.Search(id, match); err == nil {
		cli, err = httpdp.NewClient(matched, stateless)
	}
	return
}

// create a http client with IP address and port
func ClientWithAddr(addr string) (cli sbi.ConsumerClient, err error) {
	var u *url.URL
	if u, err = url.ParseRequestURI(addr); err != nil {
		return
	}
	if ep := _mesh.regcli.GetEndpoint(u.Host); ep != nil {
		cli = httpdp.NewClientWithEndpoint(ep)
	} else {
		err = fmt.Errorf("Can't find/create endpoint for %s", addr)
	}
	return
}

/*
func CallbackClient(callback string) (cli sbi.ConsumerClient, err error) {
	var addr, path string
	if addr, path, err = parseCallback(callback); err == nil {
		if ep := _mesh.regcli.GetEndpoint(addr); ep != nil {
			cli = httpdp.NewCallbackClient(ep, path)
		} else {
			err = fmt.Errorf("Can't find/create endpoint for %s", addr)
		}
	}
	return
}

// extract host address and callback prefix
func parseCallback(callback string) (host string, path string, err error) {
	var u *url.URL
	if u, err = url.ParseRequestURI(callback); err != nil {
		return
	}
	host, path = u.Host, u.Path
	return
}
*/

func AgentId() string {
	return _mesh.regcli.Id()
}

func SbiAddr() string {
	return fmt.Sprintf("%s:%d", common.GetLocalIP(), _mesh.cfg.Sbi.Port)
}

func CallbackAddress() string {
	return fmt.Sprintf("http://%s:%d", common.GetLocalIP(), _mesh.cfg.Sbi.Port)
}
