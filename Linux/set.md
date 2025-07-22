# Set Command

The `set` command in Linux is used to set or unset values of shell options and positional parameters. It can also be used to display the names and values of shell variables.

## `set -a set +a`

The `set -a` command marks all variables that are defined after it as automatically exported to the environment of subsequently executed commands. This means that any variable defined after `set -a` will be available to child processes.

The `set +a` command disables this automatic export feature, meaning that variables defined after it will not be exported to child processes.

```shell
set -a  # Enable automatic export of variables
export VAR1="value1"  # This variable will be exported
set +a  # Disable automatic export of variables
```