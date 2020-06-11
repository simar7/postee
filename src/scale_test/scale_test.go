package scale_test

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"testing"
)

const defaultCount = 10

var (
	wg sync.WaitGroup
	url string
	port string
	count int
)

func sendWebhook(image, registry int) error  {
	digest := fmt.Sprintf("%d-%d", image, registry)
	msg := buildMessage(digest, image, registry)
	r := bytes.NewReader([]byte(msg))
	resp, err := http.Post( url+":"+port, "application/json", r)
	wg.Done()
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Wrong status: %v\n", err)
	}
	return nil
}

func TestMain(m *testing.M)  {
	url = os.Getenv("TEST_URL")
	if len(url) == 0 {
		url = "http://localhost"
	}
	port = os.Getenv("TEST_PORT")
	if len(port) == 0 {
		port = "8082"
	}
	c := os.Getenv("TEST_COUNT")
	if len(c) > 0 {
		var err error
		count, err = strconv.Atoi(c)
		if err != nil {
			count = defaultCount
		}
	} else {
		count = defaultCount
	}
	m.Run()
}

func TestSendSameMessages(t *testing.T) {
	t.Logf("Start sending %d message to %s:%s", count, url, port)
	imageId := 999
	regId := 999
	for i:=0; i < count; i++ {
		wg.Add(1)
		go func(imageId, regId int) {
			err := sendWebhook(imageId, regId)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
		}(imageId, regId)
	}
	wg.Wait()
}

func TestSendUniqueMessages(t *testing.T) {
	t.Logf("Start sendting %d message to %s:%s", count, url, port)
	for i:=0; i < count; i++ {
		wg.Add(1)
		go func(imageId, regId int) {
			err := sendWebhook(imageId, regId)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
		}(i, i)
	}
	wg.Wait()
}

func buildMessage(uniqueContent string, imageId, registryId int) string {
	image := fmt.Sprintf("image%d", imageId)
	reg := fmt.Sprintf("registry%d", registryId)
	t := `{"image":"%s","registry":"%s","digest":"%s","vulnerability_summary":{"critical":0,"high":1,"medium":3,"low":4,"negligible":5},"image_assurance_results":{"disallowed":true}}`
	return fmt.Sprintf( t, image, reg, uniqueContent)
}
