# String Manipulation and Arrays

Bash provides powerful built-in mechanisms to manipulate strings and work with both indexed and associative arrays.

---

## 1. String Manipulation

### Quotes Comparison
* **Single Quotes (`'...'`)**: Literals. No variables or command substitutions are expanded.
* **Double Quotes (`"..."`)**: Expands variables (`$var`) and command substitutions (`$(...)`).

```bash
name="Kiarash"
echo 'Hello $name' # Output: Hello $name
echo "Hello $name" # Output: Hello Kiarash
```

### String Concatenation
Concatenate strings by placing variables or literals next to each other.

```bash
part1="hello"
part2="world"
greeting="$part1 $part2"
```

### Case Conversion (Bash 4+)
Modify the case of letters inside a variable.

```bash
text="Hello World"

echo "${text^^}" # Upper case -> HELLO WORLD
echo "${text,,}" # Lower case -> hello world
echo "${text^}"  # Capitalize first letter only
echo "${text,}"  # Lowercase first letter only
```

---

## 2. Indexed Arrays

Indexed arrays are standard lists ordered by numerical index (starting at 0).

### Declaration
You can define an array implicitly or explicitly.

```bash
# Implicit definition
fruits=("apple" "banana" "cherry")

# Explicit definition
declare -a sports
sports[0]="soccer"
sports[1]="basketball"
```

### Reading Elements
Access elements using indices. Wrap the variable in `${}`.

```bash
echo "${fruits[0]}" # Output: apple
echo "${fruits[-1]}" # Output: cherry (negative index matches from the end)
```

### Array Length
* `${#array[@]}`: Returns the number of elements in the array.
* `${#array[index]}`: Returns the length of the string at that specific index.

```bash
echo "Total fruits: ${#fruits[@]}" # Output: 3
echo "Length of cherry: ${#fruits[2]}" # Output: 6
```

### Iterating Over Arrays
Always quote `"${array[@]}"` to handle elements containing spaces correctly.

```bash
for fruit in "${fruits[@]}"; do
    echo "Fruit: $fruit"
done
```

### Slicing Arrays
Extract a subset of the array using `"${array[@]:offset:length}"`.

```bash
numbers=(10 20 30 40 50)
slice=("${numbers[@]:1:3}") # Captures indices 1, 2, and 3
echo "${slice[@]}"         # Output: 20 30 40
```

### Modifying Arrays
* **Append**: `array+=("new_val")`
* **Delete element**: `unset array[index]` (note: this leaves a hole in indices; see associative arrays)
* **Delete entire array**: `unset array`

```bash
fruits+=("grape")
unset fruits[1] # Removes "banana" (index 0 is "apple", index 2 is "cherry")
```

---

## 3. Associative Arrays (Key-Value Pairs - Bash 4+)

Associative arrays map unique string keys to values (often called dictionaries or maps).

> [!IMPORTANT]
> You **must** explicitly declare associative arrays using `declare -A` before using them.

### Declaration and Assignment
```bash
# Explicit declaration
declare -A user_ages

# Assignment (one-by-one)
user_ages[kiarash]=30
user_ages[sarah]=25

# Inline initialization
declare -A server_ips=(
    [web]="192.168.1.10"
    [db]="192.168.1.20"
    [cache]="192.168.1.30"
)
```

### Accessing Values
```bash
echo "DB IP: ${server_ips[db]}" # Output: 192.168.1.20
```

### Retrieving All Keys
Prefix the variable name with an exclamation mark `!`.

```bash
echo "All roles: ${!server_ips[@]}" # Output: web db cache
```

### Iterating Over Associative Arrays
Loop through keys to retrieve both keys and values.

```bash
for role in "${!server_ips[@]}"; do
    echo "The IP of $role server is ${server_ips[$role]}"
done
```
