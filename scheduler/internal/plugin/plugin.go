package plugin

var Factories = make(map[string]func(cfg Config, name string) Plugin)
var FactoriesCache = make(map[string]Plugin)

// Plugin 描述插件的接口
type Plugin interface {
	Executor(args ...string) error
	GetConfig() Config
	// Name 执行期名称
	Name() string
}

// Config 插件配置用于解析plugin.json
type Config struct {
	Name        string            `json:"name"`
	Language    string            `json:"language"`
	Suffix      string            `json:"suffix"`
	Version     string            `json:"version"`
	Description string            `json:"description"`
	Logfile     string            `json:"logfile"`
	Parameters  []ConfigParameter `json:"parameters"`
}

type ConfigParameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}
