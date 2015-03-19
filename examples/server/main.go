package main

import (
	"log"
	"os"
	"plugintest/plug/plugserver"
	"time"
)

func main() {
	ps := new(plugserver.Server)
	go ps.Run("tcp", "localhost:15235")
	go runner(ps)
	for {
		select {
		case <-ps.Quit:
			os.Exit(1)
		case <-time.After(time.Second):
			log.Print("Tick!")
		}
	}
}

// TranslationPacket is an example request/response struct
type TranslationPacket struct {
	Line1 string
	Line2 string
	Line3 string
}

func runner(ps *plugserver.Server) {
	for {
		var tx *TranslationPacket
		out := TranslationPacket{
			Line1: "Je suis la jeune fille.",
			Line2: "Zot zots the zotter?",
			Line3: "Which way did he go?",
		}
		err := ps.Process("translate", out, &tx)
		if err != nil {
			log.Print("Error: ", err)
		} else {
			if tx == nil {
				tx = &out
			}
			log.Print("WAS: ", out)
			log.Print("NOW: ", *tx)
		}
		time.Sleep(time.Second * time.Duration(5))
	}
}
