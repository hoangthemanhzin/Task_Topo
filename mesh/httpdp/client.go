package httpdp

import (
	"etrib5gc/mesh/registry"
	"etrib5gc/sbi"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BaseClient interface {
	ClientAndRemote() (*http.Client, string, error)
}
type statefulCli struct {
	*registry.Endpoint
}

func (c *statefulCli) ClientAndRemote() (*http.Client, string, error) {
	//TODO: return and error if the endpoint is not in state to send request
	return c.Client(), c.Addr(), nil
}

type statelessCli struct {
	group registry.MatchedGroup
}

func (c *statelessCli) ClientAndRemote() (*http.Client, string, error) {
	if ep, err := c.group.Select(); err != nil {
		return nil, "", err
	} else {
		return ep.Client(), ep.Addr(), nil
	}

}

type staticCli struct {
	addr string
	cli  *http.Client
}

func (c *staticCli) ClientAndRemote() (*http.Client, string, error) {
	return c.cli, c.addr, nil
}

type Client struct {
	base BaseClient
}

func NewClient(group registry.MatchedGroup, stateless bool) (cli *Client, err error) {
	if stateless {
		cli = &Client{
			base: &statelessCli{
				group: group,
			},
		}
	} else {
		var ep *registry.Endpoint
		if ep, err = group.Select(); err == nil {
			cli = &Client{
				base: &statefulCli{ep},
			}
		}
	}
	return
}
func (c *Client) DecodeResponse(rsp *sbi.Response) (err error) {
	return DecodeResponse(rsp)
}

// TODO: add re-transmission after a failure
func (c *Client) Send(req *sbi.Request) (rsp *sbi.Response, err error) {
	//send the message here
	var httpresp *http.Response
	var addr string
	var cli *http.Client

	//get client and host address
	if cli, addr, err = c.base.ClientAndRemote(); err != nil {
		return
	}
	req.Path = fmt.Sprintf("http://%s/%s", addr, req.Path)
	if err = EncodeRequest(req); err != nil {
		return
	}
	//send the request and get a response
	if httpresp, err = cli.Do(req.Request.(*http.Request)); err != nil {
		return
	} else {
		//read body of the response and prepare the sbi response
		var body []byte
		if body, err = ioutil.ReadAll(httpresp.Body); err == nil {
			rsp = &sbi.Response{
				Response:   httpresp,
				StatusCode: httpresp.StatusCode,
				Status:     httpresp.Status,
				BodyBytes:  body,
			}
		}
		httpresp.Body.Close()
	}

	return
}

func NewClientWithEndpoint(ep *registry.Endpoint) (cli *Client) {
	cli = &Client{
		base: &statefulCli{ep},
	}
	return
}

func NewClientWithAddr(addr string) (cli *Client) {
	cli = &Client{
		base: &staticCli{
			addr: addr,
			cli:  &http.Client{}, //TODO: add more settings for client
		},
	}
	return
}
