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

## `set -x set +x`

The `set -x` command enables debugging features including showing a lot more detail while running the command. `set +x` will disable it.

## `set -o set +o`

This setting prevents errors in a pipeline from being masked. If any command in a pipeline fails, that return code will be used as the return code of the whole pipeline. By default, the pipeline's return code is that of the last command even if it succeeds. 

```shell
grep some-string /non/existent/file | sort
grep: /non/existent/file: No such file or directory
echo $?
0
```

Here, grep has an exit code of 2, writes an error message to stderr, and an empty string to stdout.

This empty string is then passed through sort, which happily accepts it as valid input, and returns a status code of 0.

This is fine for a command line, but bad for a shell script: you almost certainly want the script to exit right then with a nonzero exit code... like this:

```shell
set -o pipefail
grep some-string /non/existent/file | sort
grep: /non/existent/file: No such file or directory
echo $?
2
```