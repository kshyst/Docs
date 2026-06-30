# Variables and Parameters

Variables are used to store data in memory during script execution. In Bash, variables are untyped (essentially stored as strings), but can behave as numbers or arrays depending on the context.

---

## 1. Defining and Referencing Variables

### Syntax
When assigning a value to a variable, there **must be no spaces** around the `=` operator.

```bash
# Correct
name="Kiarash"
age=30

# Incorrect (will cause commands not found errors)
name = "Kiarash"
```

### Referencing Variables
To reference a variable, prefix its name with a dollar sign `$`. You can also wrap the name in curly braces `${}` to clarify variable boundaries.

```bash
echo "Hello, $name"
echo "Your name is ${name}env" # Output: Your name is Kiarashenv
```

---

## 2. Variable Scope: Local vs. Global

By default, variables in Bash are **global** to the script (visible in functions and subshells, but not parent processes).

### Environment Variables (`export`)
Use `export` to make a variable available to child processes spawned by the script.

```bash
export API_KEY="secure-token-123"
./call_api.sh # Can access $API_KEY
```

### Local Variables (`local`)
Use the `local` keyword inside functions to restrict a variable's scope to that function.

```bash
my_function() {
    local temp_var="only visible inside here"
    echo $temp_var
}
```

---

## 3. Command Substitution

Command substitution allows you to run a command and assign its output directly to a variable.

### Modern Syntax (Recommended)
Use `$(command)`. It supports nesting and is cleaner to read.

```bash
current_dir=$(pwd)
files_count=$(ls | wc -l)
```

### Legacy Syntax
Use backticks `` `command` ``.

```bash
current_time=`date`
```

---

## 4. Special & Positional Parameters

Bash reserves special variables to handle command-line arguments and script execution state.

| Variable | Description |
| :--- | :--- |
| `$0` | The name of the script. |
| `$1` - `$9` | The first through ninth positional arguments passed to the script. |
| `${10}` ... | Arguments beyond the 9th must be wrapped in curly braces. |
| `$#` | The number of arguments passed to the script. |
| `$@` | All arguments, treated as separate quoted strings (e.g. `"$1" "$2"`). |
| `$*` | All arguments, treated as a single string (e.g. `"$1 $2"`). |
| `$?` | The exit status of the most recently executed foreground pipeline (0 for success). |
| `$$` | The Process ID (PID) of the current shell. |
| `$!` | The Process ID (PID) of the most recently executed background command. |

### Example
```bash
#!/bin/bash
echo "Script name: $0"
echo "First argument: $1"
echo "Total arguments: $#"
echo "All arguments: $@"
```

---

## 5. Advanced Parameter Expansion

Parameter expansion allows you to manipulate and modify variables on the fly.

### Default Values
* `${var:-default}`: Use `default` if `var` is unset or empty. Does not modify `var`.
* `${var:=default}`: If `var` is unset or empty, set `var` to `default` and return it.

```bash
username=${1:-"guest_user"}
```

### Length of a Variable
* `${#var}`: Returns the character length of the string inside `var`.

```bash
msg="Hello"
echo ${#msg} # Output: 5
```

### Substring Extraction
* `${var:offset}`: Extract substring from `offset` to the end.
* `${var:offset:length}`: Extract substring of `length` characters starting from `offset`.

```bash
filepath="/var/log/syslog"
echo ${filepath:5}   # Output: log/syslog
echo ${filepath:5:3} # Output: log
```

### Search and Replace
* `${var/pattern/replacement}`: Replace the first match of `pattern` with `replacement`.
* `${var//pattern/replacement}`: Replace all matches of `pattern` with `replacement`.

```bash
text="apple banana apple"
echo ${text/apple/orange}  # Output: orange banana apple
echo ${text//apple/orange} # Output: orange banana orange
```

### Strip Prefix / Suffix
* `${var#pattern}`: Remove the shortest match of `pattern` from the beginning.
* `${var##pattern}`: Remove the longest match of `pattern` from the beginning.
* `${var%pattern}`: Remove the shortest match of `pattern` from the end.
* `${var%%pattern}`: Remove the longest match of `pattern` from the end.

```bash
filename="backup.tar.gz"
echo ${filename%.*}  # Output: backup.tar (removes .gz)
echo ${filename%%.*} # Output: backup (removes .tar.gz)
```

---

## 6. Read-Only Variables

To define constants that cannot be reassigned or unset, use `readonly` or `declare -r`.

```bash
readonly DB_PORT=5432
# DB_PORT=3306 -> This will throw an error: "readonly variable"
```
