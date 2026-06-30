# Shell Scripting Overview

Shell scripting is the practice of writing a sequence of commands for a shell (command-line interpreter) to execute. It allows you to automate repetitive tasks, manage system configurations, orchestrate pipelines, and combine various command-line tools.

This section covers Shell Scripting from beginner to advanced topics.

---

## 1. The Shebang (`#!`)

A shell script is a text file that starts with a special line called the **Shebang** (`#!`). The shebang tells the operating system which interpreter to use when executing the file.

```bash
#!/bin/bash
```

Other interpreters can also be used, such as `#!/bin/sh` (POSIX compliant shell), `#!/usr/bin/env python3` (Python), or `#!/usr/bin/env node` (Node.js).

---

## 2. Making a Script Executable

By default, newly created text files do not have execute permissions. Use `chmod` to make a script executable:

```bash
chmod +x my_script.sh
```

### Execution Models
There are three main ways to run a shell script:

1. **Direct execution (uses Shebang)**:
   ```bash
   ./my_script.sh
   ```
2. **Explicit interpreter**:
   ```bash
   bash my_script.sh
   ```
3. **Sourcing (runs inside current shell context)**:
   ```bash
   source my_script.sh
   # Or:
   . my_script.sh
   ```

---

## 3. Reference Guides

Explore the detailed guides below to learn about each component of Shell Scripting:

* **[Variables & Parameters](file:///home/kshyst/Desktop/docs/docs/Linux/Shell%20Scripting/Variables.md)**: Declaring variables, global vs. local scopes, special parameter variables, and advanced parameter expansions.
* **[Conditionals & Decisions](file:///home/kshyst/Desktop/docs/docs/Linux/Shell%20Scripting/Conditions.md)**: Using standard `if` statements, comparison operators (numeric, string, file), single vs. double brackets, and `case` constructs.
* **[Loops & Control Flow](file:///home/kshyst/Desktop/docs/docs/Linux/Shell%20Scripting/Loops.md)**: Writing `for`, `while`, and `until` loops, using `break`/`continue`, and reading files line-by-line.
* **[Functions](file:///home/kshyst/Desktop/docs/docs/Linux/Shell%20Scripting/Functions.md)**: Defining reusable functions, passing parameters, variable scope, returning exit statuses or capturing string outputs.
* **[Arithmetic Operations](file:///home/kshyst/Desktop/docs/docs/Linux/Shell%20Scripting/Arithmetic.md)**: Integer math using `(( ))` and `let`, and floating-point math using the `bc` utility.
* **[Input, Output & Redirection](file:///home/kshyst/Desktop/docs/docs/Linux/Shell%20Scripting/Input_Output.md)**: Standard streams (stdin, stdout, stderr), reading user inputs, piping, redirecting streams, and Here documents.
* **[String Manipulation & Arrays](file:///home/kshyst/Desktop/docs/docs/Linux/Shell%20Scripting/Strings_Arrays.md)**: Slicing and modifying strings, case conversions, indexed arrays, and associative maps.
* **[Error Handling & Debugging](file:///home/kshyst/Desktop/docs/docs/Linux/Shell%20Scripting/Error_Handling.md)**: Understanding exit codes, using Bash strict mode (`set -euxo pipefail`), signal trapping (`trap`), and command tracing.
* **[Advanced Shell Scripting](file:///home/kshyst/Desktop/docs/docs/Linux/Shell%20Scripting/Advanced.md)**: Subshells, process substitution, process concurrency locking (`flock`), coprocesses, and indirect variables.
