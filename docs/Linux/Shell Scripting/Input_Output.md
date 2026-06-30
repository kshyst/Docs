# Input, Output, and Redirection

Managing inputs and outputs is the core of shell scripting. Programs communicate using standard streams, and Bash provides mechanisms to prompt users, format output, redirect files, and pass data between commands using pipes.

---

## 1. Standard Streams

Every process in Linux is initialized with three default open streams (Files):

| File Descriptor (FD) | Stream Name | Default Destination | Description |
| :--- | :--- | :--- | :--- |
| `0` | `stdin` | Keyboard | Standard input stream |
| `1` | `stdout` | Terminal | Standard output stream |
| `2` | `stderr` | Terminal | Standard error stream (for logs/errors) |

---

## 2. Producing Output: `echo` vs. `printf`

### `echo`
Simple and easy to use, but behavior varies between Unix systems (specifically around flag support like `-n` or `-e`).
```bash
echo "Hello World"
echo -n "No trailing newline "
echo -e "Line one\nLine two" # Enables backslash escapes
```

### `printf` (Recommended)
A robust and portable tool modeled after the C `printf` function. It gives precise control over spacing and formatting, and **does not automatically print a newline** unless specified with `\n`.

```bash
# Basic usage
printf "Hello, %s\n" "Kiarash"

# Formatting tables (Width-aligned columns)
printf "%-10s %-10s %4d\n" "Firstname" "Lastname" "Age"
printf "%-10s %-10s %4d\n" "John" "Doe" 28
```

---

## 3. Reading Input: The `read` Command

`read` pauses script execution to capture keyboard input from the user.

### Basic Prompt
```bash
read -p "Enter your username: " username
echo "Welcome $username"
```

### Silent Input (Passwords)
The `-s` flag hides user keystrokes in the terminal.
```bash
read -s -p "Enter database password: " db_password
echo "" # print newline since press-enter didn't echo it
```

### Read Options

| Option | Description | Example |
| :--- | :--- | :--- |
| `-t seconds` | Timeouts the input prompt. | `read -t 5 name` |
| `-n chars` | Exits after reading `n` characters (no need to press Enter). | `read -n 1 -p "Yes/No (y/n)?" ans` |
| `-r` | Raw mode. Prevents backslashes from acting as escape characters. | `read -r line` |

---

## 4. Redirection

Redirection allows you to send stream data to and from files instead of the terminal.

### Output Redirection
* `>`: Redirect `stdout` to a file, **overwriting** existing content.
* `>>`: Redirect `stdout` to a file, **appending** to the end of the file.

```bash
echo "First Line" > output.txt
echo "Second Line" >> output.txt
```

### Error Redirection
* `2>`: Redirect `stderr` to a file (overwrites).
* `2>>`: Redirect `stderr` to a file (appends).

```bash
# Capture errors only
ls /nonexistent-dir 2> errors.log
```

### Redirecting Both (`stdout` & `stderr`)
* `&>` or `> file 2>&1`: Redirects both output and errors to the same destination.

```bash
# Silence all output (send to the black hole /dev/null)
ping -c 1 8.8.8.8 &> /dev/null
```

### Input Redirection
* `<`: Feeds a file contents into a command's standard input.

```bash
wc -l < output.txt
```

---

## 5. Pipes (`|`)

A pipe sends the standard output (`stdout`) of the left command directly into the standard input (`stdin`) of the right command.

```bash
# Filter active processes for python
ps aux | grep "python" | wc -l
```

---

## 6. Here Documents and Here Strings

### Here Document (`<<`)
Allows you to feed multi-line text blocks directly into a command. It is commonly used to generate config files or templates.

```bash
cat << 'EOF' > /tmp/config.json
{
  "name": "App",
  "version": "1.0.0",
  "enabled": true
}
EOF
```
*Note: Quoting the delimiter (e.g. `'EOF'`) prevents variable expansion inside the block. Unquoting it allows variables to be expanded.*

### Here String (`<<<`)
Passes a single variable or string directly to a command as standard input.

```bash
message="HELLO WORLD"
# Convert string to lowercase using tr
tr 'A-Z' 'a-z' <<< "$message"
```

---

## 7. Advanced: Custom File Descriptors

You can open new file descriptors (3 through 9) to manage custom read/write streams inside your script.

```bash
# Open file descriptor 3 for writing to a log file
exec 3> "custom_event.log"

echo "Writing to custom FD 3" >&3
echo "Another message" >&3

# Close file descriptor 3
exec 3>&-
```
