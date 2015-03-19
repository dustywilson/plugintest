package main

import (
	"log"
	"plugintest/plug/plugclient"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/rsms/gotalk"
)

func main() {
	plug, err := plugclient.Connect("tcp", "localhost:15235")
	if err != nil {
		log.Fatal(err)
	}
	plug.AddHook("translate", 0, translateHook1)
	plug.AddHook("translate", 0, translateHook2)
	plug.Run()
}

// TranslationPacket is an example request/response struct
type TranslationPacket struct {
	Line1 string
	Line2 string
	Line3 string
}

func translateHook1(s *gotalk.Sock, in interface{}) (interface{}, error) {
	var tx TranslationPacket
	err := mapstructure.Decode(in, &tx)
	if err != nil {
		return in, nil
	}
	tx.Line1 = strings.Replace(tx.Line1, "fille", "girl", -1)
	log.Print(tx)
	return tx, nil
}

func translateHook2(s *gotalk.Sock, in interface{}) (interface{}, error) {
	var tx TranslationPacket
	err := mapstructure.Decode(in, &tx)
	if err != nil {
		return in, nil
	}
	tx.Line2 = strings.Replace(tx.Line2, "zot", "ZOT", -1)
	log.Print(tx)
	return tx, nil
}
