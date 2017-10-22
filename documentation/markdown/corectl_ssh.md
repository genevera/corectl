## corectl ssh

Attach to or run commands inside a running CoreOS instance

### Synopsis


Attach to or run commands inside a running CoreOS instance

```
corectl ssh VMid ["command1;..."] [flags]
```

### Examples

```
  corectl ssh VMid                 // logins into VMid
  corectl ssh VMid "some commands" // runs 'some commands' inside VMid and exits
```

### Options

```
  -h, --help          help for ssh
  -p, --port string   port to use for connecting to sshd on the VM (default "22")
```

### Options inherited from parent commands

```
  -d, --debug   adds additional verbosity, and options, directed at debugging purposes and power users
```

### SEE ALSO
* [corectl](corectl.md)	 - CoreOS over macOS made simple.

