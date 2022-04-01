package main

import (
	"fmt"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

type Step struct {
	Name       string
	UsesAction string
	Type       string
	With       map[string]string
}

type Controller struct {
	Version string
	Type    string
	Name    string
	Input   string
	Jobs    map[string]struct {
		RunsOn map[string]string
		Steps  []Step
	}
}

type Runner struct {
	Version string
	Type    string
	Config  map[string]string
}

func parseYAML(configYAML []byte, c interface{}) error {
	if err := yaml.Unmarshal(configYAML, c); err != nil {
		return err
	}
	return nil
}

func run(configController, configRunner []byte) {
	c := Controller{}
	if err := parseYAML(configController, &c); err != nil {
		panic(err)
	}

	r := Runner{}
	if err := parseYAML(configRunner, &r); err != nil {
		panic(err)
	}

	opts := &server.Options{}

	ns, err := server.NewServer(opts)
	if err != nil {
		panic(err)
	}

	go ns.Start()

	if !ns.ReadyForConnections(time.Second * 10) {
		panic("not ready for connections")
	}

	nc, err := nats.Connect(ns.ClientURL())
	if err != nil {
		panic(err)
	}

	nc.Subscribe(r.Config["name"], func(msg *nats.Msg) {
		fmt.Println(">>> data received: ", string(msg.Data))
		ns.Shutdown()
	})

	nc.Publish("postee-runner-1", []byte("hello from nats server"))
	ns.WaitForShutdown()
}
