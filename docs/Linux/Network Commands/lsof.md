# lsof

You can see the application running on specific port using:

```shell
lsof -i :5432
```

This command shows the command, PID, user running it and source and destination IP and tells of if this is a LISTENING or STABLISHED connection.

# fuser

use the fuser (file user; who uses this file) command to find all the PIDs related to that specific port. A common switch is -v which goes verbose.

```shell
sudo fuser 22/tcp -v 
```