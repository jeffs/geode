# Geode

Geode helps you manage config files for your shell, editor, etc.  You specify
settings in [TOML][] files, and Geode applies them inside a Docker container.

## Installation

First, install the prerequisites: [Git][], [Go][], and [Docker][].  Then, run
the following command:

```sh
go install github.com/jeffs/geode/...@v0.1.2
```

The `geode` command line tool should now be in your `$GOPATH/bin` directory (or
`~/go/bin` if `$GOPATH` is empty).  Re-run the command at any time to upgrade
to the latest version of Geode.  To uninstall Geode, remove the binary from
your go/bin folder, and (optionally) delete the downloaded package:

```sh
rm ~/go/bin/geode
rm -rf ~/go/pkg/*/mod/github.com/jeffs/geode
```

## Usage

Run `geode help` for a list of Geode subcommands.

## Tips

### Examples

For a trivial profile, see `testdata/groovy`.  For a larger example, see
<https://github.com/jeffs/geode-profile-home>.

### Key bindings

On the host machine, you probably want to create a `~/.docker/config.json` file
setting Docker's "detach keys" to something other than the default ^p:

```json
{ "detachKeys" : "ctrl-\\,ctrl-\\" }
```

### Clipboards

If you run containers on macOS, you can support copy/paste from the host
pasteboard via X11.  Paraphrasing [cschiewek][]:

1. Install [XQuartz](https://www.xquartz.org).
2. Under the XQuartz menu, select Preferences.
3. Under the security tab, check "Allow connections from network clients."
4. In a terminal on the host system, run `xhost + 127.0.0.1`.
5. In your Geode shell, set `DISPLAY=host.docker.internal:0`.
  - for example, in `~/.bash_profile` or `~/.zprofile`:

            export DISPLAY=host.docker.internal:0


[Docker]: https://www.docker.com/
[Git]: https://git-scm.com/
[Go]: https://golang.org/
[TOML]: https://github.com/toml-lang/toml
[cschiewek]: https://gist.github.com/cschiewek/246a244ba23da8b9f0e7b11a68bf3285
