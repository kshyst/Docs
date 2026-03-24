# vi and vim

vim is vi iMproved

In modern linux, vi  is linked to vim

## vim

#### Entring Edit Mode
- `i`: Insert mode to start edit
- `a`: Insert mode after current position of the cursor
- `o`: Enters insert mode and starts in a new line after the cursor.
- `O`: Enters insert mode and starts in a new line above the cursor.
- `esc`: Exit insert mode and enters command mode

#### Moving
- `Arrow Keys and hjkl`: Used for navigation
- `w`: Next word on the current line
- `e`: Next end of the word on the current line
- `b`: Previous beginning of the word on the current line
- `CTRL F`: Scrolls one page
- `CTRL B`: Scrolls one page backwards

#### Jumping
- `G`: Jumps to given line `23G` jumps to line 23. By default jumps to the ending line.
- `H`: Jumps to relative line from top of the screen (page). `5H`
- `L`: Jumps to relative line from bottom of the screen (page). `5L`

#### Editing

- `r`: Replaces one character
- `d`: Deletes. Use arrowkeys
- `dd`: Deletes a line
- `x`: Deletes a character at the position of the cursor
- `p`: Pastes the last deleted text after the cursor
- `P`: Pastes the last deleted text before the cursor

#### Searching
- `/`: Searches
- `?`: Searches backwards
- `n`: Repeat search (next)

#### Exiting
- `:q!`: Quit without saving
- `:w!`: Write file
- `:e!`: Reload file from disk
- `:!`: Run a shell command
- `:w newname.txt`: Write file to a new name
- `ZZ`: Exit and save if modified
- 

> Typing number before commands repeats it that many
> times. `6h` moves 6 characters to right