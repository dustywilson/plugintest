package plugclient

import "github.com/rsms/gotalk"

// Client connects to a PlugServer
type Client struct {
	sock *gotalk.Sock
}

// Connect connects a plugclient.Client to a plugserver.Server
func Connect(proto string, listen string) (*Client, error) {
	pc := new(Client)
	var err error
	pc.sock, err = gotalk.Connect(proto, listen)
	return pc, err
}

// Run runs the client
func (pc *Client) Run() {
	select {}
}
