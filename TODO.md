# To Do

* Optionally keep container alive in the absence of interactive connections.

  As a work-around, you may start a container running a do-nothing process that
  never exits.  For example:

  ```sh
  geode run PROFILE perl -MPOSIX -e pause </dev/null &
  ```

  See also: <https://unix.stackexchange.com/a/366088/49952>

* Apply configs to host system, not only inside Docker.

* Suggest custom Docker [detachKeys][], and offer to set them automatically.

  The default `CTRL-p` `CTRL-q` is too easy to hit accidentally.

  See also: <https://docs.docker.com/engine/reference/commandline/cli/>
