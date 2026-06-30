# Conditionals and Decisions

Conditional logic allows scripts to execute code selectively based on the evaluation of test expressions or exit codes.

---

## 1. The `if` Statement Syntax

The basic syntax of an `if` statement evaluates the exit status of a command. If the status is `0` (success), the code inside `then` is executed.

```bash
if command
then
    # Executed if command succeeded (returned 0)
elif command_two
then
    # Executed if command_two succeeded
else
    # Executed if all above failed
fi
```

*Note: You can place `then` on the same line as the condition by separating them with a semicolon `;`.*

---

## 2. Test Constructs: `[ ]` vs. `[[ ]]`

To evaluate expressions (like comparing strings or numbers), Bash uses test constructs.

### Single Bracket `[ expression ]` (POSIX standard)
This is an alias for the `test` command. It is highly portable but fragile.
* **Caveat**: Variables must be quoted to prevent word-splitting and globbing errors.
* **Caveat**: Uses logical operators `-a` (AND) and `-o` (OR).

```bash
if [ "$name" = "Kiarash" ]; then
    echo "Welcome"
fi
```

### Double Bracket `[[ expression ]]` (Bash Extension - Recommended)
The modern Bash-specific construct. It is much safer and more powerful.
* **Feature**: No word-splitting or globbing on variables (quoting is optional but still good practice).
* **Feature**: Direct use of `&&` and `||` for logical combinations.
* **Feature**: Support for pattern matching and regular expressions.

```bash
if [[ $name == "Kiarash" && $age -gt 18 ]]; then
    echo "Access granted"
fi
```

---

## 3. Comparison Operators

Different comparison operators are used depending on whether you are comparing **strings**, **numbers**, or **files**.

### String Comparisons
Used inside `[ ]` or `[[ ]]`. Always quote string variables to avoid errors.

| Operator | Description |
| :--- | :--- |
| `==` or `=` | True if the strings are equal. |
| `!=` | True if the strings are not equal. |
| `<` / `>` | True if string sorting is lexicographically before/after (must be escaped `\<` inside single brackets). |
| `-z string` | True if string is empty (length is 0). |
| `-n string` | True if string is not empty (length is non-zero). |

### Numeric Comparisons
Numeric comparisons are strictly integer-based.

| Operator | Description | Equivalent |
| :--- | :--- | :--- |
| `-eq` | Equal to | `==` |
| `-ne` | Not equal to | `!=` |
| `-lt` | Less than | `<` |
| `-le` | Less than or equal to | `<=` |
| `-gt` | Greater than | `>` |
| `-ge` | Greater than or equal to | `>=` |

### File Tests
Extremely useful for checking the filesystem status.

| Operator | Description |
| :--- | :--- |
| `-e file` | True if the file/directory exists. |
| `-f file` | True if it exists and is a regular file. |
| `-d file` | True if it exists and is a directory. |
| `-s file` | True if the file exists and is not empty (size > 0). |
| `-r file` | True if the file is readable. |
| `-w file` | True if the file is writable. |
| `-x file` | True if the file is executable. |
| `file1 -nt file2` | True if `file1` is newer than `file2` (modified date). |
| `file1 -ot file2` | True if `file1` is older than `file2`. |

---

## 4. Logical Operators (Combining Tests)

| Operation | Using `[ ]` | Using `[[ ]]` |
| :--- | :--- | :--- |
| **AND** | `[ "$A" = "y" -a "$B" = "z" ]` | `[[ $A == "y" && $B == "z" ]]` |
| **OR** | `[ "$A" = "y" -o "$B" = "z" ]` | `[[ $A == "y" || $B == "z" ]]` |

---

## 5. Case Statements

When matching a variable against multiple patterns, a `case` statement is cleaner than nested `if-elif` blocks.

### Syntax
```bash
case "$variable" in
    pattern1)
        # Commands
        ;;
    pattern2|pattern3)
        # Commands for pattern 2 OR pattern 3
        ;;
    *)
        # Default fallback (wildcard match)
        ;;
esac
```

### Example
```bash
read -p "Enter a file extension (e.g. sh, txt, mp3): " ext
case "$ext" in
    sh)
        echo "This is a shell script."
        ;;
    txt|md|log)
        echo "This is a text document."
        ;;
    *)
        echo "Unknown file format."
        ;;
esac
```

---

## 6. Advanced: Regular Expression Matching

With the `[[ ... ]]` syntax, you can use the `=~` operator to match a string against a Regular Expression. The matched groups are automatically captured into the `BASH_REMATCH` array.

```bash
email="user@example.com"
regex="^([a-zA-Z0-9._%+-]+)@([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$"

if [[ $email =~ $regex ]]; then
    echo "Valid email!"
    echo "Username: ${BASH_REMATCH[1]}"
    echo "Domain: ${BASH_REMATCH[2]}"
else
    echo "Invalid email format."
fi
```
