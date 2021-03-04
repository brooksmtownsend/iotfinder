# iotfinder
I have a few local devices (raspberry pis, Google coral, NVIDIA jetson) that I like to connect to with `ssh` every once in a while.

Sure I could do this with `nmap`, but I'm interested in Go and I'd like some experience with it, so here's the project.

Running this project will use `arp` to scan local addresses and test them for TCP connectivity on port 22. Currently, this is all this project is doing, but the next step will be to gather additional information about the addresses that are available for `ssh`.

## Ideal usage
My end goal for this project this is something like:
```shell
iotfinder rpi3
```
which would connect to my local raspberry pi 3 with an ssh key that I have on disk.
