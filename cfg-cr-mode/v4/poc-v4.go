package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

type Runner struct {
	nc *nats.Conn
}

func main() {
	r := Runner{}
	//defer func() {
	//	r.nc.Close()
	//}()

	// Connect Options.
	opts := setupConnOptions(nil)

	// Connect to NATS
	var err error
	r.nc, err = nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to new messages
	subjEvents := "events.postee-runner-1"
	r.nc.Subscribe(subjEvents, func(msg *nats.Msg) {
		printMsg(msg)
	})

	//subjConfig := "config.postee"
	//r.nc.Subscribe(subjConfig, func(msg *nats.Msg) {
	//printMsg(msg)
	//})

	// TODO: Periodically check for your config from the server
	//	TODO: handle the config and reconfigure the runner
	subjConfig := "config.postee"
	msg, err := r.nc.Request(subjConfig, []byte("postee-runner-1"), time.Second*5)
	fmt.Println("config: ", string(msg.Data))

	r.nc.Flush()

	log.Printf("Listening on [%s] and [%s]", subjEvents, subjConfig)

	runtime.Goexit()
}

func printMsg(m *nats.Msg) {
	log.Printf("Received on [%s]: '%s'", m.Subject, string(m.Data))
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
