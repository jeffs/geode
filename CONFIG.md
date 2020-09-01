<https://github.com/neovim/neovim/issues/3700#issuecomment-157778920>
- Windows: `~/AppData/Local`

<https://www.howtogeek.com/318177/what-is-the-appdata-folder-in-windows/>
- Windows: `${APPDATA:-~/AppData}/{Local,Roaming}`

<https://www.computerhope.com/issues/ch000109.htm>
- Windows: `$USERPROFILE` is basically `~`

<https://golang.org/pkg/os/#UserHomeDir>
- GoLang: `os.UserHomeDir()` returns `~` per `$HOME` or `$USERPROFILE`

<https://stackoverflow.com/a/5084892/3116635>
- `XDG_CONFIG_HOME`: `~/Library/Preferences/`
- `XDG_DATA_HOME`: `~/Library/`
- `XDG_CACHE_HOME`: `~/Library/Caches/`

<https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html>
- `XDG_CACHE_HOME`: `~/.cache`
- `XDG_CONFIG_DIRS`: `/etc/xdg`
- `XDG_CONFIG_HOME`: `~/.config`
- `XDG_DATA_DIRS`: `/usr/local/share:/usr/share`
- `XDG_DATA_HOME`: `~/.local/share`
- `XDG_RUNTIME_DIR`: basically a RAM disk for fifos and small files

Dirs are ordered; first value wins.  HOME implicitly precedes DIRS.

`*_HOME` and `XDG_RUNTIME_DIR` indicate writable locations, whereas `*_DIRS`
represent search paths implicitly preceded by the corresponding `*_HOME` paths.
