package outputs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aquasecurity/postee/v2/data"
	"github.com/aquasecurity/postee/v2/formatting"
	"github.com/aquasecurity/postee/v2/layout"
	"github.com/aquasecurity/postee/v2/log"
)

type WebhookOutput struct {
	Name string
	Url  string
}

func (webhook *WebhookOutput) GetName() string {
	return webhook.Name
}

func (webhook *WebhookOutput) CloneSettings() *data.OutputSettings {
	return &data.OutputSettings{
		Name:   webhook.Name,
		Url:    webhook.Url,
		Enable: true,
		Type:   "webhook",
	}
}

func (webhook *WebhookOutput) Init() error {
	log.Logger.Infof("Starting Webhook output %q, for sending to %q",
		webhook.Name, webhook.Url)
	return nil
}

func (webhook *WebhookOutput) Send(content map[string]string) error {
	log.Logger.Infof("Sending webhook to %q", webhook.Url)
	data := content["description"] //it's not supposed to work with legacy renderer
	resp, err := http.Post(webhook.Url, "application/json", strings.NewReader(data))
	if err != nil {
		log.Logger.Errorf("Sending webhook Error: %v", err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Errorf("Sending %q Error: %v", webhook.Name, err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		msg := "sending webhook wrong status: %q. Body: %s"
		return fmt.Errorf(msg, resp.StatusCode, body)
	}
	log.Logger.Infof("Sending Webhook to %q was successful!", webhook.Name)
	return nil
}

func (webhook *WebhookOutput) Terminate() error {
	log.Logger.Infof("Webhook output %q terminated.", webhook.Name)
	return nil
}

func (webhook *WebhookOutput) GetLayoutProvider() layout.LayoutProvider {
	// Todo: This is MOCK. Because Formatting isn't need for Webhook
	// todo: The App should work with `return nil`
	return new(formatting.HtmlProvider)
}
