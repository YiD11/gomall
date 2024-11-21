package mq

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(os.Getenv("NATS_ADDRESS"))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if err != nil {
		log.Panicln(err)
	}

}
