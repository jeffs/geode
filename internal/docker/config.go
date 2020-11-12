package docker

import (
	"path"
	"strings"

	"github.com/BurntSushi/toml"
)

type config struct {
	Base string
	Name string
	User string

	Locale   string
	TimeZone string

	Command []string

	Bind    map[string]string
	Volumes map[string]string
	Ports   map[string]int
}

// FlatName returns a valid Docker container name based on c.Name.  In
// particular, whereas image names may contain slashes (/), container names
// cannot.
func (c *config) FlatName() string {
	return strings.ReplaceAll(c.Name, "/", "-")
}

func readConfig(profile string) (*config, error) {
	var c config
	f := path.Join(profile, "docker.toml")
	if _, err := toml.DecodeFile(f, &c); err != nil {
		return nil, err
	}

	if c.Name == "" {
		_, c.Name = path.Split(strings.TrimRight(profile, "/"))
	}

	return &c, nil
}
