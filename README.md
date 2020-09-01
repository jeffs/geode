# Geode

Geode helps you manage your config files for your shell, editor, etc.  You
specify settings in [TOML][] files, and Geode applies those to your host system
or a Docker container.

## Installation

First, install the prerequisites: [Git][], [Go][], and optionally [Docker][]
(if you want Geode to create containers).  Then, run the following command:

    go get github.com/jeffs/geode

The `geode` command line tool should now be in your `$GOPATH/bin` directory, or
`~/go/bin` if `$GOPATH` is empty.  Finally, run `geode update` to complete the
installation.

[Docker]: https://www.docker.com/
[Git]: https://git-scm.com/
[Go]: https://golang.org/
[TOML]: https://github.com/toml-lang/toml
