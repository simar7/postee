package scale_test

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"testing"
)

var (
	wg sync.WaitGroup
)

func sendWebhook(image, registry int) error  {
	digest := fmt.Sprintf("%d-%d", image, registry)
	msg := buildMessage(digest, image, registry)
	r := bytes.NewReader([]byte(msg))
	_, err := http.Post("http://localhost:8082", "application/json", r)
	wg.Done()
	return err
}

func TestSendSameMessages(t *testing.T) {
	const count = 5
	imageId := 27
	regId := 25
	for i:=0; i < count; i++ {
		wg.Add(1)
		go sendWebhook(imageId, regId)
	}
	wg.Wait()
}

func TestSendUniqueMessages(t *testing.T) {
	const count = 1000
	for i:=0; i < count; i++ {
		wg.Add(1)
		go sendWebhook(i, i)
	}
	wg.Wait()
}

func buildMessage(uniqueContent string, imageId, registryId int) string {
	image := fmt.Sprintf("image%d", imageId)
	reg := fmt.Sprintf("registry%d", registryId)
	t := `{"image":"%s","registry":"%s","digest":"%s","vulnerability_summary":{"critical":0,"high":1,"medium":3,"low":4,"negligible":5},"image_assurance_results":{"disallowed":true}}`
	return fmt.Sprintf( t, image, reg, uniqueContent)
}
