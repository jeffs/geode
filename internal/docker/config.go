package docker

import (
	"log"
	"os/user"
	"path"
	"path/filepath"
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
	Init    []string

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

func expandGuest(orig, home string) string {
	if path.IsAbs(orig) {
		return orig
	}

	// Use / as the path separator, regardless of host platform.
	return path.Join(home, orig)
}

func expandHost(orig, home string) string {
	if filepath.IsAbs(orig) {
		return orig
	}

	// Use host-appropriate path separator.
	return filepath.Join(home, orig)
}

// Expand relative paths in Bind (keys and values).
func expandBindPaths(bind map[string]string, guestHome, hostHome string) map[string]string {
	r := make(map[string]string)
	for k, v := range bind {
		r[expandHost(k, hostHome)] = expandGuest(v, guestHome)
	}

	return r
}

// Expand relative paths in Volumes (values only).
func expandVolumePaths(volumes map[string]string, guestHome string) {
	for k, v := range volumes {
		volumes[k] = expandGuest(v, guestHome)
	}
}

func readConfig(profile string) (*config, error) {
	var c config
	f := path.Join(profile, "docker.toml")
	if _, err := toml.DecodeFile(f, &c); err != nil {
		return nil, err
	}

	// By default, use the profile directory name for the image name.
	if c.Name == "" {
		c.Name = filepath.Base(strings.TrimRight(profile, "/"))
	}

	// By default, run a do-nothing init script.
	if len(c.Init) < 1 {
		c.Init = []string{"true"}
	}

	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	g := filepath.Join("/home/", c.User) // in the guest (container)
	h := u.HomeDir                       // on the host machine
	c.Bind = expandBindPaths(c.Bind, g, h)
	expandVolumePaths(c.Volumes, g)

	return &c, nil
}
