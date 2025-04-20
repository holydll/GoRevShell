# GoRevShell
A small reverse shell written in go
this reverse shell works on Linux and macOS, for Windows, replace `/bin/sh` with `cmd.exe` and remove the `Setsid` attribute.

It uses TCP to connect back to the attacker's IP and sends shell input/output over the socket.
