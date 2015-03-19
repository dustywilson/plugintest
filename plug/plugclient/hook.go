package plugclient

import (
	"log"
	"plugintest/plug/common"
)

// AddHook calls the server and adds a plugin hook
func (pc *Client) AddHook(topic string, priority int, pf PluginFunc) {
	var response common.HookResponse
	pc.sock.Request("AddHook", common.HookRequest{
		Name:     topic,
		Priority: priority,
	}, &response)
	log.Print(response)
	pc.sock.Handlers.Handle(response.RequestName, pf)
}
