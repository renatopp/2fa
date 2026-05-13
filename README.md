# 2FA Authenticator

## Usage

First thing to do is adding a new entry in the keyring: 

```
2fa add <name> <secret_key>
```

If you don't provide a secret key, you'll be prompt.

With an entry registered, you just need to consult it:

```
2fa show <name>
```

## Installation

Install using: 

```
go install github.com/renatopp/2fa@latest
```

The OS X implementation depends on the /usr/bin/security binary for interfacing with the OS X keychain. It should be available by default. The Linux and BSD implementation depends on the Secret Service dbus interface, which is provided by GNOME Keyring and should also be available by default. If gnome-keyring is not available, for example in WSL, you can install by:

```
sudo apt install gnome-keyring
``` 

It's expected that the default collection `login` exists in the keyring, because it's the default in most distros. If it doesn't exist, you can create it through the keyring frontend program Seahorse:

```
sudo apt install seahorse
```

Then:

- Open seahorse
 - Go to File > New > Password Keyring
- Click Continue
- When asked for a name, use: login