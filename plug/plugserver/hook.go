package plugserver

import (
	"errors"
	"log"
	"plugintest/plug/common"

	"code.google.com/p/go-uuid/uuid"

	"github.com/mitchellh/mapstructure"
	"github.com/rsms/gotalk"
)

// Process sends the payload to the registered hook plugins
func (ps Server) Process(hook string, value interface{}, out interface{}) error {
	ps.hooksLock.RLock()
	defer ps.hooksLock.RUnlock()
	if len(ps.hooks[hook]) == 0 {
		return nil
	}
	for _, plugin := range ps.hooks[hook] {
		plugin.sock.Request(plugin.requestName, value, out)
		if err := mapstructure.Decode(out, &value); err != nil {
			return err
		}
	}
	return nil
}

func (ps Server) addHook(s *gotalk.Sock, request common.HookRequest) (*common.HookResponse, error) {
	hook := request.Name
	priority := request.Priority

	if hook == "" {
		return nil, errors.New("bad hook name")
	}

	ps.hooksLock.Lock()
	defer ps.hooksLock.Unlock()

	if _, ok := ps.hooks[hook]; !ok {
		ps.hooks[hook] = make(map[int]*Plugin)
	}

	for {
		if _, ok := ps.hooks[hook][priority]; ok {
			priority++
		} else {
			break
		}
	}

	uuid := uuid.New()

	plugin := &Plugin{
		sock:        s,
		hook:        hook,
		priority:    priority,
		requestName: uuid,
	}
	ps.hooks[hook][priority] = plugin
	log.Print(plugin)

	return &common.HookResponse{
		Name:        hook,
		Priority:    priority,
		RequestName: uuid,
	}, nil

	log.Print(1)
	peer := ps.peerFromSock(s)
	log.Print(2)
	peer.Lock() // FIXME: lock problem here?
	log.Print(3)
	defer peer.Unlock()
	log.Print(4)
	peer.addPlugin(plugin)
	log.Print(5)
	log.Print(peer)

	return nil, nil
}
