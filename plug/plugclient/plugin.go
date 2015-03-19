package plugclient

import "github.com/rsms/gotalk"

// PluginFunc handles plugin hooks on the client side
type PluginFunc func(*gotalk.Sock, interface{}) (interface{}, error)
