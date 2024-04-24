package golang

import (
	"github.com/ideal-rucksack/workflow-scheduler/pkg/consotants/cfg"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/plugin"
	"os"
	"os/exec"
	"path/filepath"
)

type Plugin struct {
	Cfg        plugin.Config
	PluginName string
}

func (p *Plugin) Executor(args ...string) error {
	config := p.GetConfig()
	pluginPath := filepath.Join(os.Getenv(cfg.PluginHome), p.PluginName, config.Name)
	err := exec.Command(pluginPath, "-webhook", "http://localhost:5266/webhooks", "-action", "databases").Run()
	return err
}

func NewGolangPlugin(cfg plugin.Config, name string) plugin.Plugin {
	return &Plugin{Cfg: cfg, PluginName: name}
}

func (p *Plugin) GetConfig() plugin.Config {
	return p.Cfg
}

func (p *Plugin) Name() string {
	return p.PluginName
}
