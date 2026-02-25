| Name       | Function                                                                                          |
|------------|---------------------------------------------------------------------------------------------------|
| `USER`     | The name of the logged-in user                                                                    |
| `PATH`     | List of directories to search for commands, colon separated                                       |
| `EDITOR`   | Default editor                                                                                    |
| `HISTFILE` | Where bash should save its history (normally .bash_history)                                       |
| `HOSTNAME` | System hostname                                                                                   |
| `PS1`      | The Prompt! Play with it                                                                          |
| `UID`      | The numeric user id of the logged-in user                                                         |
| `HOME`     | The user's home directory                                                                         |
| `PWD`      | The current working directory                                                                     |
| `SHELL`    | The name of the shell                                                                             |
| `$`        | The process id (or PID) of the running bash shell (or other) process                             |
| `PPID`     | The process id of the process that started this process (that is, the id of the parent process)   |
| `?`        | The exit code of the last command                                                                 |

>Global bash configs are stored at `/etc/profile` and each user has their config at `~/.profile` & `~/.bash\_profile` & `~/.bash\_logout`. If you need a permanent change, add your configs to these.