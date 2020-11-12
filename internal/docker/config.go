package docker

import (
	"path"

	"github.com/BurntSushi/toml"
)

type config struct {
	Base string
	Name string
	User string

	Locale   string
	TimeZone string

	Command []string
}

func readConfig(profile string) (*config, error) {
	var c config
	f := path.Join(profile, "dockerfile.toml")
	if _, err := toml.DecodeFile(f, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
