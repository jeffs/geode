# To Do

* Optionally keep container alive in the absence of interactive connections.

  As a work-around, you may start a container running a do-nothing process that
  never exits.  For example:

  ```sh
  geode run PROFILE perl -MPOSIX -e pause </dev/null &
  ```

  See also: <https://unix.stackexchange.com/a/366088/49952>
