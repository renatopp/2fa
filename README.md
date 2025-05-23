# 2FA Authenticator

OS X
The OS X implementation depends on the /usr/bin/security binary for interfacing with the OS X keychain. It should be available by default.

Linux and *BSD
The Linux and *BSD implementation depends on the Secret Service dbus interface, which is provided by GNOME Keyring.

It's expected that the default collection login exists in the keyring, because it's the default in most distros. If it doesn't exist, you can create it through the keyring frontend program Seahorse:

sudo apt install gnome-keyring

Open seahorse
Go to File > New > Password Keyring
Click Continue
When asked for a name, use: login