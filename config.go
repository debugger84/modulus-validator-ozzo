package validator

type ModuleConfig struct {
}

func NewModuleConfig() *ModuleConfig {
	return &ModuleConfig{}
}
func (s *ModuleConfig) ProvidedServices() []interface{} {
	return []interface{}{
		func() *ModuleConfig {
			return s
		},
	}
}
