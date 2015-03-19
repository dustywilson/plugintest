package plugserver

import "github.com/rsms/gotalk"

// Plugin is a representation of the connected plugin
type Plugin struct {
	sock        *gotalk.Sock
	hook        string
	priority    int
	requestName string
}
