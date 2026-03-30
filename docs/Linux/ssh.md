# Encryption

## ssh-keygen

Stores key under `~/.ssh`. System wide keys are under `/etc/ssh`

- `-t`: Chooses algorithm. Default is `rsa`

Algorithms:
`dsa`, `ecdsa`, `ecdsa-sk`, `ed25519`, `ed25519-sk`, `rsa`

When the fingerprint of the server changes we get this:

```shell
ssh 192.168.70.2
Host key verification failed.
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
IT IS POSSIBLE THAT SOMEONE IS DOING SOMETHING NASTY!
Someone could be eavesdropping on you right now (man-in-the-middle attack)!
It is also possible that a host key has just been changed.
The fingerprint for the ED25519 key sent by the remote host is
SHA256:4Wp2zz6sgPAhnbqhkNjOd6QDpNQ4jvjX7qAzslPX09U.
Please contact your system administrator.
Add correct host key in /Users/jadi/.ssh/known_hosts to get rid of this message.
Offending RSA key in /Users/jadi/.ssh/known_hosts:548
Host key for 192.168.70.2 has changed and you have requested strict checking.
```

To fix it use :

```shell
ssh-keygen -R 192.168.70.2
```

## Passwordless Login

- Add public key to `~/.ssh/authorized_keys` 
- `/etc/ssh/sshd_config` should contain `PubkeyAuthentication yes`

To move the public key automatically:

```shell
ssh-copy-id 192.168.70.2
```

## ssh-agent

The ssh-agent works like a key manager for your ssh keys. It keeps your private keys in memory and provides them when needed.

For example if you have set a password on your key, you have to type the password every single time you are going to use that key. To make this easier using ssh-agent, you have to rush a shell with ssh-agent and add all your key to the agent:

```shell
ssh-agent /bin/bash
ssh-add
```
Now, the agent knows about your keys and wont ask for the passwords anymore.

`ssh-agent` can also transfer your keys safely in memory to other servers. Say you need to ssh to server A and while on server A, use your key to ssh to server B (or do a git pull using your private keys). The keys are on server A and its not safe to actually and physically copy them to server B. In this case you cal run ssh-agent and have your keys even when you are on server B.


## ssh-tunnel

### Local Forwarding

```shell
ssh -L 9000:hckrnews.com:80 root@5.161.197.79
curl localhost:9000
```

Here I'm telling my computer to ssh using user `root` 
(which is not a great idea) to the `5.161.197.79` server; 
AND create a `local tunnel (-L)` from on my machine toward 
the `hckrnews.com port 80` through that machine. Now if I 
connect to `localhost:9000` on my machine, the request will 
be tunnels through `5.161.197.79` toward `hckrnews.com port 80`.
You can try it with `curl localhost:9000`.

Why this is useful? Say you have a program on your server which only answers back to the local requests (and not the internet). Using local forwarding you can forward a port on your computer to port which programs works on and use the program on your own machine!

### Remote Forwarding

There is also the concept of Remote Forwarding. In this case, you connect a port from a remote server to another server (mostly your own computer). Using this, you can expose a webserver from your own computer on the Internet:

```shell
ssh -R 8000:localhost:80 root@5.161.197.79
```

In the above example, I'm telling the ssh to create a Remote tunnel. After this any request on port 8000 from the 5.161.197.79 computer will reach to localhost (my machine) port 80!

You can even ask the remote machine to start listening on a specific port on all of its interfaces (open up to the network and not only localhost):

```shell
ssh -R 0.0.0.0:8000:localhost:7777 192.168.70.2
```

Above, I'm telling the 192.168.70.2 machine to open port 8000 to ALL interfaces and forward whatever it got to my machiens 7777 port.

> to let clients to bind listening ports to anything other than localhost, the GatewayPorts clientspecified configu should be set in the ssh server's config file.

### Dynamic Application Level Forwarding

use the -D switch for a dynamic application level port forwarding. It works like a proxy / anti censorship.

```shell
ssh -D 1080 192.168.70.2
```

The 1080 port will work ask a socks proxy on my machine and forward whatever request reaches it to the 192.168.70.2 machine and returns back the answers. Now I can configure my applications to use the localhost:1080 as their socks proxy and the 192.168.70.2 will work as a proxy here. This is useful when you do not have internet on your localhost, but 192.168.70.2 has it.

### X Forwarding

`X11Forwarding yes` in `sshd_config`

```shell
ssh -X 192.168.70.2
```