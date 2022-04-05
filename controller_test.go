//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//
//	"github.com/stretchr/testify/require"
//)
//
//func TestParseYaml(t *testing.T) {
//	b, err := ioutil.ReadFile("cfg-cr-mode/cfg-controller-v4.yaml")
//	require.NoError(t, err)
//
//	c := Controller{}
//	parseYAML(b, &c)
//	assert.Equal(t, "foo", c.Jobs["Kill bad process, send logs job"])
//
//	b, err = ioutil.ReadFile("cfg-cr-mode/cfg-runner-v4.yaml")
//	r := Runner{}
//	parseYAML(b, &r)
//	fmt.Println(r)
//	assert.Equal(t, "foo", r.Config)
//}
//
//func TestRun(t *testing.T) {
//	c, err := ioutil.ReadFile("cfg-cr-mode/cfg-controller-v4.yaml")
//	require.NoError(t, err)
//	r, err := ioutil.ReadFile("cfg-cr-mode/cfg-runner-v4.yaml")
//	require.NoError(t, err)
//
//	run(c, r)
//}
