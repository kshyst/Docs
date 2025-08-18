# Internet Message Access Protocol(IMAP)

IMAP allows synchronized read, move and delete messages. Mainly used to check emails from multiple clients. POP3 reads and deletes mail from remote server. IMAP keep emails and synchronize them across multiple clients.

Some of the IMAP Commands:

- `LOGIN <username> <password>` authenticates the user
- `SELECT <mailbox>` selects the mailbox folder to work with
- `FETCH <mail_number> <data_item_name>` Example fetch 3 body[] to fetch message number 3, header and body.
- `MOVE <sequence_set> <mailbox>` moves the specified messages to another mailbox
- `COPY <sequence_set> <data_item_name>` copies the specified messages to another mailbox
- `LOGOUT` logs out

IMAP listens on port 143 by default