# Geode

Geode helps you manage your config files for your shell, editor, etc.  You
specify settings in [TOML][] files, and Geode applies those to your host system
or a Docker container.

## Installation

First, install the prerequisites: [Git][], [Go][], and optionally [Docker][]
(if you want Geode to create containers).  Then, run the following command:

```sh
go get github.com/jeffs/geode
```

The `geode` command line tool should now be in your `$GOPATH/bin` directory (or
`~/go/bin` if `$GOPATH` is empty).

## Usage

Run `geode help` for a list of Geode subcommands.

## Tips

On the host machine, you probably want to create a `~/.docker/config.json` file
setting Docker's "detach keys" to something other than the default ^p:

```json
{ "detachKeys" : "ctrl-\\,ctrl-\\" }
```

[Docker]: https://www.docker.com/
[Git]: https://git-scm.com/
[Go]: https://golang.org/
[TOML]: https://github.com/toml-lang/toml
