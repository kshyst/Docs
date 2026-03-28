# About Nginx

Traditionally, web servers (like Apache) created a new “process” or “thread” for every single user request. If 10,000 people visited, the server needed 10,000 threads, which ate up massive amounts of RAM.

NGINX uses an event-driven, asynchronous architecture. Instead of creating a new process for every user, a single NGINX worker process can handle thousands of concurrent connections. This makes it incredibly fast and lightweight.