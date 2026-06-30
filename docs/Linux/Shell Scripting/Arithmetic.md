# Arithmetic Operations

Bash treats all variables as strings by default. To perform mathematical calculations, you must use specific operators and constructs. Bash natively supports integer arithmetic; floating-point arithmetic requires external utilities like `bc`.

---

## 1. The Double Parentheses Construct `(( ))`

The double parentheses construct is the most modern, clean, and recommended way to perform integer math in Bash.

### A. Assigning Results
To capture the result of a calculation into a variable, prefix it with a dollar sign `$(( expression ))`.

```bash
sum=$((5 + 3))
product=$((sum * 4))
echo "Result: $product" # Output: 32
```

### B. Inline Evaluation / Side-Effects
You can perform math directly inside `(( ))` without the prefix `$`. This is useful for counter increments and modifications.

```bash
counter=0
((counter++))      # Increment counter by 1
((counter += 5))   # Add 5 to counter
echo $counter      # Output: 6
```

### C. Logic and Comparisons
`(( ))` can also be used as a test condition in `if` statements, allowing direct use of standard comparison symbols (`<`, `>`, `<=`, `>=`, `==`, `!=`).

```bash
score=85
if ((score >= 80)); then
    echo "Excellent!"
fi
```

---

## 2. Legacy Methods

### The `let` Command
`let` parses arithmetic expressions directly. It is equivalent to inline `(( ))` but requires quotes if spaces are present.

```bash
let "val = 10 * 3"
let val++
```

### The `expr` Command (Legacy POSIX)
An external utility that prints results to standard output. It is slow, requires spaces between all operands, and requires escaping operators like `*`.

```bash
result=$(expr 5 + 3)
product=$(expr 5 \* 3)
```

---

## 3. Supported Arithmetic Operators

Inside `(( ))` or `$(( ))`, you can use standard programming operators:

| Operator | Description | Example |
| :--- | :--- | :--- |
| `+` | Addition | `$((10 + 5))` |
| `-` | Subtraction | `$((10 - 5))` |
| `*` | Multiplication | `$((10 * 5))` |
| `/` | Division (Integer division, drops remainder) | `$((10 / 3))` -> `3` |
| `%` | Modulo (Remainder of division) | `$((10 % 3))` -> `1` |
| `**` | Exponentiation | `$((2 ** 3))` -> `8` |

### Bitwise and Shift Operators
For low-level binary manipulation:
* Shift left / right: `<<` / `>>`
* Bitwise AND / OR / XOR: `&` / `|` / `^`
* Bitwise NOT: `~`

---

## 4. Advanced: Floating-Point Math with `bc`

Because Bash only supports integer arithmetic, division like `5 / 2` will return `2`. For decimal and floating-point math, pipe the expression to `bc` (Basic Calculator).

### Usage
Pipe a string containing the calculation into `bc`. Use `scale` to specify the number of decimal places.

```bash
# Division with 2 decimal places
result=$(echo "scale=2; 5 / 2" | bc)
echo "Result: $result" # Output: 2.50
```

### Complex Math Functions
Pass the `-l` option to `bc` to load the math library (which enables sine, cosine, logs, exponential, etc.).

```bash
# Calculate square root
root=$(echo "scale=4; sqrt(10)" | bc)
echo "Square root of 10: $root" # Output: 3.1622
```

### Floating-Point Comparisons
You can also use `bc` to evaluate truth conditions for decimals.

```bash
val1=2.5
val2=1.8

# bc returns 1 if true, 0 if false
if [[ $(echo "$val1 > $val2" | bc) -eq 1 ]]; then
    echo "$val1 is larger than $val2"
fi
```
