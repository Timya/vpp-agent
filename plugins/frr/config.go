package frr

type Config struct {
	frrVersion string `json:"frr-version"`
}

func DefaultConfig() *Config {
	return &Config{
		frrVersion: "test-version-frr",
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
		p.Log.Debugf("config file %q not found, using default configls", p.Cfg.GetConfigName())
	}
	return cfg, nil
}
