# Geode

Geode helps you manage config files for your shell, editor, etc.  You specify
settings in [TOML][] files, and Geode applies them inside a Docker container.

## Installation

First, install the prerequisites: [Git][], [Go][], and [Docker][].  Then, run
the following command:

```sh
go get github.com/jeffs/geode/...
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

For a trivial profile, see `testdata/groovy`.  For a larger example, see
<https://github.com/jeffs/geode-profile-home>.

[Docker]: https://www.docker.com/
[Git]: https://git-scm.com/
[Go]: https://golang.org/
[TOML]: https://github.com/toml-lang/toml
