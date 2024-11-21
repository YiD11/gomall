package notify

import (
	"log"

	email "github.com/YiD11/gomall/rpc_gen/kitex_gen/email"
	"github.com/kr/pretty"
)

type NoopEmail struct {
}

// simulate sending email
func (n *NoopEmail) Send(req *email.EmailReq) error {
	pretty.Printf("%v\n", req)
	log.Printf("%+v\n", req)
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
