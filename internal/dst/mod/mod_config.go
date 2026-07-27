package mod

import "gihora/pkg/util"

const (
	DstConfigPath = "pkg/config/dst_config.yaml"
)

type Config struct {
	ModsSetupPath string `yaml:"mods_setup_path"`
}

func InitConfig() (Config, error) {
	var cfg Config
	if err := util.ReadYAML(DstConfigPath, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
