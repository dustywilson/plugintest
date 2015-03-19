package plugserver

import (
	"sync"

	"github.com/rsms/gotalk"
)

// Peer is the connected plugin provider
type Peer struct {
	sock    *gotalk.Sock
	plugins []*Plugin
	*sync.RWMutex
}

type peerFromSockRequest struct {
	sock     *gotalk.Sock
	respChan chan *Peer
}

func (ps *Server) newPeer(s *gotalk.Sock) *Peer {
	return &Peer{
		sock:    s,
		plugins: make([]*Plugin, 0),
	}
}

func (ps *Server) peerFromSock(s *gotalk.Sock) *Peer {
	respChan := make(chan *Peer)
	ps.peerFromSockChan <- peerFromSockRequest{
		sock:     s,
		respChan: respChan,
	}
	return <-respChan
}

func (pr *Peer) addPlugin(plugin *Plugin) {
	pr.plugins = append(pr.plugins, plugin)
}

// Quit cleans up the disconnected peer
func (pr *Peer) Quit() {
}
