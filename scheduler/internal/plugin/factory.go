package plugin

type FactoryRegistry struct {
	factories map[string]Plugin
}

func NewFactoryRegistry() *FactoryRegistry {
	return &FactoryRegistry{
		factories: make(map[string]Plugin),
	}
}

func (r *FactoryRegistry) Register(name string, plugin Plugin) {
	r.factories[name] = plugin
}

func (r *FactoryRegistry) GetPlugin(name string) (Plugin, bool) {
	plugin, ok := r.factories[name]
	if !ok {
		return nil, false
	}
	return plugin, true
}
