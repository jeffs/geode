# Simplest usage

    git clone github.com:jeffs/geode.git
    go run ./install.go
      => creates "${XDG_CONFIG_HOME:-$HOME/.config}"/geode

    geode install
    vim PROFILE.toml
    geode build PROFILE

geode build PROFILE.toml
  => puts 
  

geode build PROFILE  => IMAGE
geode exec  IMAGE
Run the installer (in the same repo)

## [XDG][] is crap
It covers only "user specific" files.
It uses the misspelling "separated."

[XDG]: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
