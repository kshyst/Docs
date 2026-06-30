# Advanced Shell Scripting

This section covers advanced Bash scripting concepts, including subshells, process substitution, process concurrency, coprocesses, dynamic referencing, and signal management.

---

## 1. Subshells

A subshell is a child process spawned by the current shell. It inherits the environment variables and file descriptors of the parent shell, but operates in an isolated environment.

### Syntax
Wrap commands in parentheses `( )`.

```bash
message="Global value"

(
    # Inside subshell
    message="Subshell value"
    echo "Subshell: $message"
)

echo "Parent: $message" # Output: Global value
```

> [!NOTE]
> Changes made to variables, current directories (`cd`), or settings inside a subshell do not affect the parent shell.

---

## 2. Process Substitution

Process substitution allows you to treat the output of a command as if it were a temporary file path. It is denoted by `<(command)` (for reading) and `>(command)` (for writing).

### A. Difference from Pipes
Normally, a pipe runs commands in a subshell, preventing variable mutations from persisting.

```bash
# Variables modified in the pipe loop are lost
count=0
cat items.txt | while read -r line; do
    ((count++))
done
echo "Total: $count" # Output: 0

# Using Process Substitution resolves this
count=0
while read -r line; do
    ((count++))
done < <(cat items.txt)
echo "Total: $count" # Output: (Actual count)
```

### B. Comparing Command Outputs
You can pass multiple command outputs directly to utilities like `diff` without writing to temporary files on disk.

```bash
diff <(ssh server1 "ls /var/log") <(ssh server2 "ls /var/log")
```

---

## 3. Concurrency and Locking (`flock`)

When running scripts via cron jobs, you may want to prevent a script from executing if a previous instance is still running. You can use the `flock` utility to manage script locks.

### Inline Lock Script
```bash
#!/bin/bash
exec 9<"$0" # Open current script file for reading on FD 9
flock -n 9 || { echo "Script is already running!"; exit 1; }

# Critical section starts here
echo "Running critical tasks..."
sleep 10
```

---

## 4. Indirect Variable References

Indirect referencing allows you to retrieve the value of a variable whose name is stored inside another variable.

### Syntax (Bash 2.0+)
Use the `${!var}` format.

```bash
real_var="SecureToken123"
pointer="real_var"

echo "${!pointer}" # Output: SecureToken123
```

---

## 5. Coprocesses (`coproc`)

A coproc starts a command in the background, creating two-way pipe communication between the parent shell and the background command.

### Example
```bash
# Start background coprocess named 'math_engine'
coproc math_engine { bc; }

# Write calculation to the coprocess input stream
echo "5 * 5" >&"${math_engine[1]}"

# Read result from the coprocess output stream
read -r answer <&"${math_engine[0]}"

echo "Result: $answer" # Output: 25

# Terminate coprocess
kill $math_engine_PID
```

---

## 6. Trap for Signal Management

In advanced scripting, you may want to handle system signals (like SIGHUP or SIGUSR1) to perform tasks like reloading configurations dynamically without restarting the script process.

```bash
#!/bin/bash
echo "Process ID: $$"

reload_config() {
    echo "Reloading configuration file..."
    # Read config commands here
}

# Trap SIGUSR1 signal (signal number 10)
trap reload_config SIGUSR1

# Infinite loop simulating daemon
while true; do
    sleep 10
done
```

*Note: You can trigger this trap from another terminal using `kill -10 <PID>` or `kill -USR1 <PID>`.*
