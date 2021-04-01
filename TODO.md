# To Do

* GUI support is currently terrible
  - Enabling GUIs ought to be simple; e.g., a flag or docker.toml property
  - GUIs (via XQuartz) are prohibitively slow
* Optionally keep containers alive in the absence of interactive connections
  - As a work-around, start a container running an immortal do-nothing process
  - For example, call [pause(2)][], as in [premount.zsh][]
  - `geode run PROFILE perl -MPOSIX -e pause </dev/null &`
* Apply configs to host system, not only inside Docker
* Suggest custom Docker [detachKeys][], and offer to set them automatically
  - The default `CTRL-p` `CTRL-q` is too easy to hit accidentally
  - See also the [Docker CLI reference][]


[Docker CLI reference]: https://docs.docker.com/engine/reference/commandline/cli/
[pause(2)]: https://unix.stackexchange.com/a/366088/49952
[premount.zsh]: https://github.com/jeffs/geode-profile-home/blob/master/premount.zsh
