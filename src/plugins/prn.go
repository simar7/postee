package plugins

import (
	"formatting"
	"layout"
	"log"
	"settings"
)

type PrnPlugin struct {
	PrnSetting *settings.Settings
	provider   layout.LayoutProvider
}

func (prn *PrnPlugin) Init() error {
	prn.provider = new(formatting.HtmlProvider)
	log.Printf("Starting Prn plugin %q...", prn.PrnSetting.PluginName)
	return nil
}

func (prn *PrnPlugin) Send(input map[string]string) error {
	log.Printf("Sending via %q: %q", prn.PrnSetting.PluginName, input["title"])
	return nil
}

func (prn *PrnPlugin) Terminate() error {
	return nil
}

func (prn *PrnPlugin) GetLayoutProvider() layout.LayoutProvider {
	return prn.provider
}

func (prn *PrnPlugin) GetSettings() *settings.Settings {
	return prn.PrnSetting
}