# Daemon

In multitasking computer operating systems, a daemon is a computer program that runs as a background process, rather than being under the direct control of an interactive user.

## Daemon Processes

In multitasking computer operating systems, a daemon is a computer program that runs as a background process, rather than being under the direct control of an interactive user.\

## How Daemon Processes Are Created

A daemon process is created by detaching a process from its parent and associating it with the init process. This is done in several steps: first, the process calls fork() to create a child process. Then, the child process calls setsid() to create a new session and detach from any terminal. Afterward, it changes its working directory to root (/), resets file permissions, and redirects standard input, output, and error to /dev/null. Finally, it runs in the background to perform its tasks.

![Daemon Process](https://media.geeksforgeeks.org/wp-content/uploads/20250115122617728744/daemon.webp)

for further information, you can check [this link](https://www.geeksforgeeks.org/daemon-processes/)