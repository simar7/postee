package v5

import (
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"

	"github.com/goccy/go-yaml"
)

type Job struct {
	Name       string
	UsesAction string
	Type       string
	With       map[string]string
}

type RunnerConfig struct {
	Name  string
	Input string
	Jobs  []Job
}

type Controller struct {
	Version       string
	Type          string
	Name          string
	RunnerConfigs []RunnerConfig `yaml:"runner-configs"`

	server   *server.Server
	conn     *nats.Conn
	shutdown chan struct{}
}

type Runner struct {
	Version string
	Type    string
	Config  map[string]string

	conn *nats.Conn
}

func parseYAML(configYAML []byte, c interface{}) error {
	if err := yaml.Unmarshal(configYAML, c); err != nil {
		return err
	}
	return nil
}

//func (c *Controller) Init(configController []byte) {
//	if err := parseYAML(configController, c); err != nil {
//		panic(err)
//	}
//}
//
//func (c *Controller) Run() {
//	opts := &server.Options{Host: "0.0.0.0", Port: 8888}
//
//	ns, err := server.NewServer(opts)
//	if err != nil {
//		panic(err)
//	}
//
//	go ns.Start()
//
//	if !ns.ReadyForConnections(time.Second * 10) {
//		panic("not ready for connections")
//	}
//
//	c.conn, err = nats.Connect(ns.ClientURL())
//	if err != nil {
//		panic(err)
//	}
//	//nc.Publish("postee-runner-1", []byte("hello from nats server"))
//
//	select {
//	case <-c.shutdown:
//		ns.Shutdown()
//	}
//
//	//<-c.shutdown
//	//ns.Shutdown()
//}
//
//func (c Controller) Send() {
//	c.conn.Publish("postee-runner-1", []byte("hello from nats server"))
//}
//
//func (r *Runner) Init(configRunner []byte) {
//	if err := parseYAML(configRunner, r); err != nil {
//		panic(err)
//	}
//}
//
//func (r Runner) Run() {
//	nc, err := nats.Connect(r.Config["postee-controller-url"])
//	if err != nil {
//		panic(err)
//	}
//
//	nc.Subscribe(r.Config["name"], func(msg *nats.Msg) {
//		fmt.Println(">>> data received: ", string(msg.Data))
//		//ns.Shutdown()
//	})
//}
//
