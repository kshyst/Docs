# SSH

To ssh a server without providing password:

Generate ssh key on local machine and give the public key to the server.

```shell
ssh-keygen
```

- Private key → `~/.ssh/id_rsa`
- Public key → `~/.ssh/id_rsa.pub`

Copying the public key to server:

```shell
ssh-copy-id user@server_ip
```

Or just manually put it in :

```shell
~/.ssh/authorized_keys
```