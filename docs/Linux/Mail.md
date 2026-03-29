# Mail Transfer Agents (MTA)

Email is an integral part of many GNU/Linux and Unix systems. Each user do have a mail box and can send / receive email to other local users. This is done via MTAs (Mai Transfer Agents). In other words, MTAs are programs which handle emails in your operating system. They can receive and dispatch emails locally and over the network. There are different options for MTAs. In this section we will do a quick review on them and you will see how you can send emails to other uses (or over the internet) and how you can check your local mails.

### sendmail
Is one of the oldest options available. It is huge and kind of difficult to configure and not that security oriented. Because of these, few systems use it as default their MTA.

### exim
It aims to be a general and flexible mailer with extensive facilities for checking incoming e-mail. It is feature rich with ACLs, authentication and many other features.

### postfix
This is a new alternative to sendmail and uses easy to understand configuration files. It supports multiple domains, encryption, etc. Postfix is what you see on most distros as the default MTA.

> Most desktop distros do not install MTAs by default. If you want, I suggest installing the postfix (and `mailx` or `bsd-mailx)` via your package manager.


As you already know, `sendmail` is the oldest MTA alive and therefore, many other MTAs try to comply with it and has a `sendmail emulation layer` to keep themselves backward compatible with sendmail. Thats why you can type sendmail on whatever distro you are or use the mailq and check your mail regardless of your MTA choice.


### aliases

There are some mail aliases on the system. Defined in `/etc/aliases`.

```shell
postmaster:    root
```

Means if mail comes for postmaster, send it to root.

In case of any change in this file, you need to run the `newaliases` command.

Other way: Each user can create a `.forward` file in her own directory and all mail targeted to that user will be forwarded to that address.

> You can even put a complete email address like jadijadi@gmail.com in your .forward file.

## Sending Mail

Use `mail` command

```shell
mail postmaster
```

After finishing press `Ctrl+D`. 

## mailq

This command lists the mail queue. Each entry shows the queue file ID, message size, arrival time, sender, and the recipients that still need to be delivered. If mail could not be delivered upon the last attempt, the reason for failure is shown. The sysadmin can use this command to check the status of emails still in the queues.