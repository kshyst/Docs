# Post Office Protocol version 3

Designed to listen to the mail servers and receive emails.

Common POP3 commands:

- `USER <username>` identifies the user
- `PASS <password>` provides the user’s password
- `STAT` requests the number of messages and total size
- `LIST` lists all messages and their sizes
- `RETR` <message_number> retrieves the specified message
- `DELE` <message_number> marks a message for deletion
- `QUIT` ends the POP3 session applying changes, such as deletions

POP3 Listens on port 110 TCP by default.

