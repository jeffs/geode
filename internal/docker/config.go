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

	Bind map[string]string
	Volumes map[string]string
	Ports map[string]int
}

func readConfig(profile string) (*config, error) {
	var c config
	f := path.Join(profile, "docker.toml")
	if _, err := toml.DecodeFile(f, &c); err != nil {
		return nil, err
	}

	// TODO: Populate Name (if not explicitly set) from profile path.

	return &c, nil
}
