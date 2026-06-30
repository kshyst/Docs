# Error Handling and Debugging

Writing robust shell scripts requires anticipating failures, handling unexpected return codes, cleaning up resources, and using debugging features to trace issues.

---

## 1. Exit Codes

Every command in Linux returns an exit status (an integer between `0` and `255`) when it terminates.

* **`0`**: Success.
* **Non-zero (e.g. `1` to `255`)**: Failure.

### Checking Exit Code (`$?`)
The special variable `$?` holds the exit code of the most recently executed foreground command.

```bash
mkdir /root/secret_folder
if [[ $? -ne 0 ]]; then
    echo "Failed to create directory. Do you have root privileges?"
    exit 1
fi
```

### Common Exit Codes

| Exit Code | Meaning | Example / Cause |
| :--- | :--- | :--- |
| `1` | General error | Miscellaneous errors, e.g. divide by zero. |
| `2` | Misuse of shell builtins | Syntax error or missing arguments in a shell builtin. |
| `126` | Command invoked cannot execute | Permission problem or command is not executable. |
| `127` | Command not found | typo in command name, or path not in `$PATH`. |
| `128` | Invalid argument to `exit` | `exit 3.14` (only integers allowed). |
| `128 + N` | Fatal error signal `N` | Script terminated by signal `N` (e.g. `130` for SIGINT / Ctrl+C). |

---

## 2. Short-Circuiting Error Checks

Instead of writing full `if` statements, you can use logical operators `&&` (AND) and `||` (OR) to handle errors inline.

```bash
# Execute command_two only if command_one succeeded
command_one && command_two

# Exit the script immediately if cd fails
cd /var/log || exit 1
```

---

## 3. Bash Strict Mode (`set` options)

By default, Bash scripts will keep running even if a command fails or if they reference a variable that doesn't exist. You can enforce strict error policies using the `set` builtin.

```bash
# Combine them for the "Unofficial Bash Strict Mode"
set -euxo pipefail
```

### Explaining the Options:

#### `set -e` (`errexit`)
Exit immediately if any command exits with a non-zero status.
* **Exception**: Commands evaluated inside conditionals (e.g. `if command; then`) do not trigger an exit.

#### `set -u` (`nounset`)
Treat unset variables as an error and exit immediately. Prevents bugs like accidentally deleting directories due to an empty path variable (`rm -rf "$MIPATH/*"`).

#### `set -o pipefail`
Normally, a pipeline's exit status is determined solely by the last command (e.g. `fail_command | success_command` returns `0`). `pipefail` ensures the pipeline returns the exit code of the last (rightmost) command that failed.

#### `set -x` (`xtrace`)
Prints each command to standard error before running it. Excellent for tracing execution.

---

## 4. Cleaning Up with `trap`

Scripts often create temporary files or start background jobs. The `trap` command allows you to catch signals (like exit or termination) and execute clean-up code before the script terminates.

### Syntax
```bash
trap 'cleanup_commands' SIGNAL_NAMES
```

### Common Signals
* `EXIT`: Fires when the script finishes (successfully or due to errors/`set -e`).
* `SIGINT` (Ctrl+C): Fires when interrupted by the user.
* `SIGTERM`: Fires when a kill signal is received.

### Example
```bash
#!/bin/bash
set -e

# Create a temporary file
TEMP_FILE=$(mktemp)
echo "Temp file is at $TEMP_FILE"

# Define cleanup function to delete the file
cleanup() {
    echo "Cleaning up temp files..."
    rm -f "$TEMP_FILE"
}

# Register the trap for EXIT and SIGINT
trap cleanup EXIT SIGINT

# Rest of the script
echo "Performing work..."
sleep 5
```

---

## 5. Debugging Tools

### Syntax Check
To parse a script for syntax errors without executing it, use the `-n` flag.

```bash
bash -n my_script.sh
```

### Dry Run with Output
To trace script execution step-by-step, run with the `-x` flag.

```bash
bash -x my_script.sh
```
