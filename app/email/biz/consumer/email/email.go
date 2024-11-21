package email

import (
	"context"
	"fmt"
	"log"

	"github.com/YiD11/gomall/app/email/infra/mq"
	"github.com/YiD11/gomall/app/email/infra/notify"
	email "github.com/YiD11/gomall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

func InitConsumer() {
	tracer := otel.Tracer("shop-nats-consumer")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	sub, err := mq.Nc.Subscribe("email", func(m *nats.Msg) {
		var req email.EmailReq
		fmt.Println(":recevive ", string(m.Data))
		err := proto.Unmarshal(m.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}

		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(m.Header))	
		_, span := tracer.Start(ctx, "shop-nats-consumer")

		defer span.End()

		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})
	if err != nil {
		log.Panicln(err)
	}

	server.RegisterStartHook(func() {
		sub.Unsubscribe() // nolint:errcheck
		mq.Nc.Close()
	})

}
