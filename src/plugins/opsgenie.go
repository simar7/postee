package plugins

import (
	"layout"
	"settings"
)

type OpsGeniePlugin struct {
	OpsGenieSettings *settings.Settings
	opsGenieLayout   layout.LayoutProvider
}

func (ops *OpsGeniePlugin) Init() error {
	return nil
}
func (ops *OpsGeniePlugin) Send(map[string]string) error { return nil }
func (ops *OpsGeniePlugin) Terminate() error             { return nil }
func (ops *OpsGeniePlugin) GetLayoutProvider() layout.LayoutProvider {
	return ops.opsGenieLayout
}
func (ops *OpsGeniePlugin) GetSettings() *settings.Settings {
	return ops.OpsGenieSettings
}
