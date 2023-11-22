package ngap

import (
	"encoding/hex"
	"etrib5gc/nfs/pran/context"
	"etrib5gc/sctp"
	"fmt"
	"io"
	"net"
	"sync"
	"syscall"

	//	"git.cs.nctu.edu.tw/calee/sctp"

	libngap "github.com/free5gc/ngap"
)

type Server struct {
	connections sync.Map
	listener    *sctp.SCTPListener
	ngap        *Ngap
	wg          sync.WaitGroup
	iplist      []string
	port        int
	done        chan bool
}

const readBufSize uint32 = 8192

// set default read timeout to 2 seconds
var readTimeout syscall.Timeval = syscall.Timeval{Sec: 2, Usec: 0}

var sctpConfig sctp.SocketConfig = sctp.SocketConfig{
	InitMsg:   sctp.InitMsg{NumOstreams: 3, MaxInstreams: 5, MaxAttempts: 2, MaxInitTimeout: 2},
	RtoInfo:   &sctp.RtoInfo{SrtoAssocID: 0, SrtoInitial: 500, SrtoMax: 1500, StroMin: 100},
	AssocInfo: &sctp.AssocInfo{AsocMaxRxt: 4},
}

func NewServer(ctx *context.CuContext, iplist []string, port int) *Server {
	_initLog()
	return &Server{
		ngap:   NewNgap(ctx),
		done:   make(chan bool),
		iplist: iplist,
		port:   port,
	}
}

func (s *Server) Run() (err error) {

	if s.listener != nil {
		err = fmt.Errorf("Server is running")
		log.Errorf(err.Error())
		return
	}

	ips := []net.IPAddr{}
	var netAddr *net.IPAddr
	for _, addr := range s.iplist {
		if netAddr, err = net.ResolveIPAddr("ip", addr); err != nil {
			log.Errorf("Error resolving address '%s': %v\n", addr, err)
			return
		} else {
			log.Infof("Resolved address '%s' to %s\n", addr, netAddr)
			ips = append(ips, *netAddr)
		}
	}

	addr := &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    s.port,
	}

	if s.listener, err = sctpConfig.Listen("sctp", addr); err != nil {
		log.Errorf("Failed to listen: %+v", err)
		return
	}
	log.Infof("Listen on %s", s.listener.Addr())
	go s.loop()
	return
}

func (s *Server) loop() {
	s.wg.Add(1)
	defer s.wg.Done()
	log.Infof("Start waiting for Ngap connections")
	for {
		select {
		case <-s.done:
			return
		default:
			newConn, err := s.listener.AcceptSCTP()
			if err != nil {
				log.Errorf("a error :%s", err.Error())
				switch err {
				case syscall.EINTR, syscall.EAGAIN:
					log.Errorf("AcceptSCTP: %+v", err)
				default:
					log.Errorf("Failed to accept: %+v", err)
				}
				continue
			} else if newConn != nil {
				log.Infof("a new connection from %s:%s", newConn.LocalAddr().Network(), newConn.LocalAddr().String())
				var info *sctp.SndRcvInfo
				if infoTmp, err := newConn.GetDefaultSentParam(); err != nil {
					log.Errorf("Get default sent param error: %+v, accept failed", err)
					if err = newConn.Close(); err != nil {
						log.Errorf("Close error: %+v", err)
					}
					continue
				} else {
					info = infoTmp
					log.Debugf("Get default sent param[value: %+v]", info)
				}

				info.PPID = libngap.PPID
				if err := newConn.SetDefaultSentParam(info); err != nil {
					log.Errorf("Set default sent param error: %+v, accept failed", err)
					if err = newConn.Close(); err != nil {
						log.Errorf("Close error: %+v", err)
					}
					continue
				} else {
					log.Debugf("Set default sent param[value: %+v]", info)
				}

				events := sctp.SCTP_EVENT_DATA_IO | sctp.SCTP_EVENT_SHUTDOWN | sctp.SCTP_EVENT_ASSOCIATION
				if err := newConn.SubscribeEvents(events); err != nil {
					log.Errorf("Failed to accept: %+v", err)
					if err = newConn.Close(); err != nil {
						log.Errorf("Close error: %+v", err)
					}
					continue
				} else {
					log.Debugln("Subscribe SCTP event[DATA_IO, SHUTDOWN_EVENT, ASSOCIATION_CHANGE]")
				}

				if err := newConn.SetReadBuffer(int(readBufSize)); err != nil {
					log.Errorf("Set read buffer error: %+v, accept failed", err)
					if err = newConn.Close(); err != nil {
						log.Errorf("Close error: %+v", err)
					}
					continue
				} else {
					log.Debugf("Set read buffer to %d bytes", readBufSize)
				}

				if err := newConn.SetReadTimeout(readTimeout); err != nil {
					log.Errorf("Set read timeout error: %+v, accept failed", err)
					if err = newConn.Close(); err != nil {
						log.Errorf("Close error: %+v", err)
					}
					continue
				} else {
					log.Debugf("Set read timeout: %+v", readTimeout)
				}

				log.Infof("SCTP Accept from: %s", newConn.RemoteAddr().String())
				s.connections.Store(newConn, newConn)

				go s.handleConnection(newConn, readBufSize)
			} else {
				log.Trace("Listener timeouted")
			}
		}
	}
}

func (s *Server) Stop() {
	if s.listener == nil {
		log.Infof("No SCTP server to stop")
		return
	}

	log.Infof("Close SCTP server...")
	if err := s.listener.Close(); err != nil {
		log.Error(err)
		log.Infof("SCTP server may not close normally.")
	}
	close(s.done)
	log.Infof("Close connections...")
	s.connections.Range(func(key, value interface{}) bool {
		conn := value.(net.Conn)
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
		return true
	})
	s.wg.Wait()
	log.Infof("SCTP server closed")
}

func (s *Server) handleConnection(conn *sctp.SCTPConn, bufsize uint32) {
	s.wg.Add(1)
	defer func() {
		// if Pran call Stop(), then conn.Close() will return EBADF because
		// conn has been already closed
		if err := conn.Close(); err != nil && err != syscall.EBADF {
			log.Errorf("close connection error: %+v", err)
		}
		log.Infof("close/delete connection ")
		s.connections.Delete(conn)
		s.wg.Done()
	}()

	buf := make([]byte, bufsize)
	for {
		select {
		case <-s.done:
			return
		default:
			n, info, notification, err := conn.SCTPRead(buf)
			if err != nil {
				switch err {
				case io.EOF, io.ErrUnexpectedEOF:
					log.Errorf("Read EOF from client")
					return
				case syscall.EAGAIN:
					log.Debug("SCTP read timeout")
					continue
				case syscall.EINTR:
					log.Errorf("SCTPRead: %+v", err)
					continue
				default:
					log.Errorf("Handle connection[addr: %+v] error: %+v", conn.RemoteAddr(), err)
					return
				}
			}

			if notification != nil {
				s.ngap.HandleSCTPNotification(conn, notification)
			} else {
				if info == nil || info.PPID != libngap.PPID {
					log.Warnln("Received SCTP PPID != 60, discard this packet")
					continue
				}

				log.Tracef("Read %d bytes", n)
				log.Tracef("Packet content:\n%+v", hex.Dump(buf[:n]))

				// TODO: concurrent on per-UE message
				s.ngap.HandleMessage(conn, buf[:n])
			}
		}
	}
}
