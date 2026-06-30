# Functions

Functions allow you to group commands into reusable blocks of code. They reduce duplication, make scripts modular, and improve readability.

---

## 1. Defining Functions

In Bash, you can define functions in two ways. The POSIX standard syntax is preferred for portability.

### POSIX Standard Syntax (Recommended)
```bash
my_function() {
    # commands
    echo "This is standard syntax"
}
```

### Alternative Syntax
```bash
function my_function {
    # commands
    echo "This is alternative syntax"
}
```

### Invoking Functions
To invoke a function, write its name as you would any other shell command. **Do not use parentheses** when calling it.

```bash
my_function # Correct call
```

---

## 2. Function Arguments

Functions do not declare parameters in their signatures. Instead, arguments passed to a function are accessed inside the function using positional parameters (`$1`, `$2`, `$#`, `$@`).

> [!IMPORTANT]
> Within a function, positional parameters `$1`, `$2`, etc. refer to the **function's arguments**, not the command-line arguments passed to the script.

```bash
greet() {
    echo "Hello, $1! You are $2 years old."
    echo "Number of args passed to function: $#"
}

# Invoking the function with arguments
greet "Kiarash" 30
```

---

## 3. Variable Scope: Global vs. Local

By default, all variables in Bash are global. If you define or modify a variable inside a function, it will affect the entire script.

### Using the `local` Keyword
To prevent side effects, always declare variables inside functions with the `local` keyword. This limits their scope to the function block.

```bash
global_var="I am global"

my_func() {
    local local_var="I am local to my_func"
    global_var="Modified by function!"
    echo "Inside: $local_var"
    echo "Inside: $global_var"
}

my_func
echo "Outside: $global_var"  # Output: Modified by function!
# echo $local_var            # Output: (nothing, it's out of scope)
```

---

## 4. Returning Values from Functions

Bash functions do not return objects or strings in the traditional programming sense. They can return value in two main ways.

### A. Return Status Code (`return`)
The `return` statement exits the function and sets the return status code (an integer between `0` and `255`). By convention, `0` indicates success, and any non-zero value indicates an error.

```bash
check_file() {
    if [[ -f "$1" ]]; then
        return 0 # Success
    else
        return 1 # Failure
    }
}

check_file "/etc/passwd"
if [[ $? -eq 0 ]]; then
    echo "File exists!"
fi
```

### B. Returning Data / Output (Recommended)
To return strings, arrays, or numbers to the main script, send the output to standard output (`stdout`) and capture it using command substitution.

```bash
get_domain() {
    local email="$1"
    # Print the domain to stdout
    echo "${email##*@}"
}

# Capture stdout into a variable
domain=$(get_domain "kiarash@example.com")
echo "Domain is: $domain" # Output: example.com
```

---

## 5. Exporting Functions

If your script calls a subshell or another script, your functions are not automatically visible to them. You can export a function using `export -f`.

```bash
log_message() {
    echo "[LOG] $1"
}

export -f log_message

# Spawn a subshell or execute a separate script that uses log_message
bash -c 'log_message "Hello from subshell"'
```

---

## 6. Advanced: Recursive Functions

Bash supports recursion (functions that call themselves). You must ensure there is a base case to terminate execution.

```bash
# Calculate factorial
factorial() {
    local num=$1
    if ((num <= 1)); then
        echo 1
    else
        local prev=$(factorial $((num - 1)))
        echo $((num * prev))
    fi
}

result=$(factorial 5)
echo "Factorial of 5 is: $result" # Output: 120
```
