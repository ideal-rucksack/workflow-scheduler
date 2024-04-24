package plugin

import (
	"encoding/json"
	"errors"
	"fmt"
	plugin2 "github.com/ideal-rucksack/workflow-scheduler/pkg/consotants/plugin"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/logging"
	"os"
	"path/filepath"
)

// Loader 插件加载器
type Loader interface {
	// Load 加载插件
	Load(name string) (Plugin, error)
}

// DefaultPluginLoader 默认插件加载器
type DefaultPluginLoader struct {
	Path string
}

// Load 加载插件 未来也许我们要支持互联网下载插件
func (d DefaultPluginLoader) Load(name string) (Plugin, error) {
	var (
		err    error
		plugin Plugin
		cfg    Config
	)

	if name == "" {
		return nil, errors.New("plugin name is empty")
	} else {
		plugin, ok := FactoriesCache[name]
		if ok {
			return plugin, nil
		}
	}

	if d.Path == "" {
		return nil, errors.New("plugin path cannot be empty")
	}

	_, err = os.Stat(d.Path)

	if err != nil {
		return nil, err
	}

	pluginHome, err := os.Stat(filepath.Join(d.Path, name))

	if err != nil {
		return nil, err
	}

	if !pluginHome.IsDir() {
		return nil, fmt.Errorf("plugin %s does not exist", name)
	}

	// 校验插件配置
	pluginCfg, err := os.Open(filepath.Join(d.Path, name, plugin2.CfgFilename+"."+plugin2.CfgFiletype))
	if err != nil {
		return nil, fmt.Errorf("failed to open plugin config: %w", err)
	}

	defer func(pluginCfg *os.File) {
		err := pluginCfg.Close()
		if err != nil {
			logging.Logger.Errorf("failed to close plugin config file: %s", err)
		}
	}(pluginCfg)

	err = json.NewDecoder(pluginCfg).Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode plugin config: %w", err)
	}

	factoryFunc, ok := Factories[cfg.Language]
	if !ok {
		return nil, fmt.Errorf("unsupported plugin language: %s", cfg.Language)
	}

	plugin = factoryFunc(cfg, name)
	FactoriesCache[name] = plugin

	logging.Logger.Infof("plugin %s loaded", name)

	return plugin, nil
}
