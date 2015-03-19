package plugserver

import (
	"log"
	"sync"

	"github.com/rsms/gotalk"
)

// Server is the system to handle plugins
type Server struct {
	gotalk           *gotalk.Server
	hooks            map[string]map[int]*Plugin
	hooksLock        sync.RWMutex
	peers            map[*gotalk.Sock]*Peer
	connected        chan *gotalk.Sock
	disconnected     chan *gotalk.Sock
	peerFromSockChan chan peerFromSockRequest
	Quit             chan bool
}

// Run runs the Plug Server
func (ps *Server) Run(proto string, listen string) error {
	ps.Quit = make(chan bool)
	ps.hooks = make(map[string]map[int]*Plugin)
	ps.peers = make(map[*gotalk.Sock]*Peer)
	ps.connected = make(chan *gotalk.Sock)
	ps.disconnected = make(chan *gotalk.Sock)
	ps.peerFromSockChan = make(chan peerFromSockRequest)

	var err error
	ps.gotalk, err = gotalk.Listen(proto, listen)
	if err != nil {
		close(ps.Quit)
		return err
	}

	ps.gotalk.AcceptHandler = ps.pluginConnected
	ps.gotalk.Handlers.Handle("AddHook", ps.addHook)

	go ps.connectionManager()
	err = ps.gotalk.Accept()

	close(ps.Quit)
	return err
}

func (ps *Server) connectionManager() {
	for {
		select {
		case <-ps.Quit:
			return
		case s := <-ps.connected:
			ps.peers[s] = ps.newPeer(s)
		case s := <-ps.disconnected:
			peer := ps.peers[s]
			delete(ps.peers, s)
			peer.Quit()
		case r := <-ps.peerFromSockChan:
			log.Print("PFSC 1")
			r.respChan <- ps.peers[r.sock]
			log.Print("PFSC 2")
		}
	}
}

func (ps *Server) pluginConnected(s *gotalk.Sock) {
	s.CloseHandler = func(s *gotalk.Sock, code int) {
		ps.pluginDisconnected(s, code)
	}
	ps.connected <- s
}

func (ps *Server) pluginDisconnected(s *gotalk.Sock, code int) {
	ps.disconnected <- s
}
