package v5

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseYaml(t *testing.T) {
	b, err := ioutil.ReadFile("configs/cfg-controller-v5.yaml")
	require.NoError(t, err)

	c := Controller{}
	require.NoError(t, parseYAML(b, &c))
	assert.Equal(t, "Send slack message", c.RunnerConfigs[0].Jobs[0].Name)

	b, err = ioutil.ReadFile("configs/cfg-runner-v5.yaml")
	r := Runner{}
	require.NoError(t, parseYAML(b, &r))
	assert.Equal(t, "postee-runner-1", r.Config["name"])
}

func TestRun(t *testing.T) {
	cCfg, err := ioutil.ReadFile("configs/cfg-controller-v5.yaml")
	require.NoError(t, err)
	rCfg, err := ioutil.ReadFile("configs/cfg-runner-v5.yaml")
	require.NoError(t, err)

	c := Controller{}
	require.NoError(t, parseYAML(cCfg, &c))

	r := Runner{}
	require.NoError(t, parseYAML(rCfg, &r))

	c.server = test.RunRandClientPortServer()
	//c.server, err = server.NewServer(&server.Options{Debug: true, TraceVerbose: true, Trace: true})
	//c.server, err = server.NewServer(&server.Options{Host: "0.0.0.0", Port: 9999})
	require.NoError(t, err)
	go c.server.Start()

	if !c.server.ReadyForConnections(10 * time.Second) {
		t.Fatal("not ready for connections")
	}

	c.conn, err = nats.Connect(c.server.ClientURL()) // controller connection to publish
	require.NoError(t, err)

	r.conn, err = nats.Connect(c.server.ClientURL()) // runner connection to subscribe
	require.NoError(t, err)

	_, err = r.conn.Subscribe("postee-runner-1", func(msg *nats.Msg) {
		fmt.Println(">>> data received: ", string(msg.Data))
		c.server.Shutdown()
	})
	require.NoError(t, err)
	fmt.Println("subbed")
	time.Sleep(time.Millisecond)

	err = c.conn.Publish("postee-runner-1", []byte("hello from server"))
	require.NoError(t, err)
	fmt.Println("pubbed")

	c.server.WaitForShutdown()
}
