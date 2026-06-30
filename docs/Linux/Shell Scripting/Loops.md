# Loops and Control Flow

Loops allow you to execute a block of code multiple times. Bash supports `for`, `while`, and `until` loops, along with loop control statements like `break` and `continue`.

---

## 1. The `for` Loop

The `for` loop is used to iterate over a list of items or a sequence of values.

### A. List-Based Loop
Iterates through a static list of strings or the output of a command/glob.

```bash
for fruit in apple banana orange; do
    echo "I like $fruit"
done
```

### B. Range-Based Sequence Loop
Uses brace expansion `{start..end..step}` to generate a sequence.

```bash
# Count from 1 to 5
for i in {1..5}; do
    echo "Number: $i"
done

# Count from 0 to 10 by twos
for i in {0..10..2}; do
    echo "Even: $i"
done
```

### C. C-Style `for` Loop
Uses double parentheses `(( ))` for numeric loop conditions.

```bash
for ((i=0; i<5; i++)); do
    echo "Index: $i"
done
```

### D. Iterating Over File Globs (Safely)
Always wrap variables in quotes inside the loop body to handle filenames with spaces correctly.

```bash
# Loop through all markdown files
for file in *.md; do
    # Prevent running if no files match the glob
    [[ -e "$file" ]] || continue
    echo "Processing document: $file"
done
```

---

## 2. The `while` Loop

The `while` loop executes a block of code repeatedly as long as the test condition evaluates to true (returns `0`).

```bash
counter=1
while [[ $counter -le 5 ]]; do
    echo "Count: $counter"
    let counter=counter+1 # Or: ((counter++))
done
```

### Infinite Loops
An infinite loop runs indefinitely unless terminated by a `break` statement or an external signal.

```bash
while true; do
    echo "Press [CTRL+C] to stop..."
    sleep 2
done
```

---

## 3. The `until` Loop

The `until` loop is the opposite of the `while` loop. It continues executing the block of code as long as the test condition is false, terminating when the condition becomes true.

```bash
counter=1
until [[ $counter -gt 5 ]]; do
    echo "Count: $counter"
    ((counter++))
done
```

---

## 4. Loop Control: `break` and `continue`

* **`break`**: Immediately exits the loop.
* **`continue`**: Skips the rest of the current iteration and starts the next loop cycle.

```bash
for num in {1..10}; do
    # Skip odd numbers
    if ((num % 2 != 0)); then
        continue
    fi
    
    # Stop the loop if number exceeds 6
    if ((num > 6)); then
        break
    fi
    
    echo "Valid Even Number: $num"
done
```

---

## 5. Advanced: Reading Files and Inputs Line-by-Line

The most robust way to process text files or command outputs line-by-line is combining a `while` loop with the `read` command.

### Reading From a File
Using input redirection `<`.

```bash
# -r prevents backslash escaping from being processed
while read -r line; do
    echo "Line content: $line"
done < "users.txt"
```

### Reading From Command Output
Using a pipe `|`.

```bash
# Note: This runs the loop inside a subshell, so variables modified inside
# the loop will not persist outside of it.
df -h | grep "/dev/" | while read -r filesystem size used avail percent mount; do
    echo "Mountpoint $mount is at $percent capacity"
done
```

### Process Substitution (Preserves Variable State)
If you need to read a file line-by-line and modify variables that must persist outside the loop, use process substitution instead of a pipe.

```bash
total_lines=0
while read -r line; do
    ((total_lines++))
done < <(grep "ERROR" system.log)

echo "Total errors found: $total_lines"
```
