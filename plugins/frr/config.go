package frr

type Config struct {
	FrrVersion string `json:"frr-version"`
}

func DefaultConfig() *Config {
	return &Config{
		FrrVersion: "test-version-frr",
	}
}

func (p *FRRPlugin) loadConfig() (*Config, error) {

	cfg := DefaultConfig()
	found, err := p.Cfg.LoadValue(cfg)

	if err != nil {
		return nil, err
	} else if found {
		p.Log.Debugf("config loaded from file %q", p.Cfg.GetConfigName())
	} else {
		p.Log.Infof("config file %q not found, using default config: %q", p.String(), p.Cfg.GetConfigName())

	}
	return cfg, nil
}
