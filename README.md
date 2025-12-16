# Haiku Go Tools

This is a fork of `golang.org/x/tools` to make some tools compatible 
with the Go 1.18 shipped with Haiku OS repositories.

Currently it contains only [`github.com/aligator/haiku-golang-tools/gopls`](https://pkg.go.dev/github.com/aligator/haiku-golang-tools/gopls).

If something other of `golang.org/x/tools` needs to be made compatible with Haiku, 
please open an issue and I may look into it.

For further documentation refer to the official documentation.

## Install Gopls

Basically all from https://go.dev/gopls/ applies except that you need to install it using

```bash
go install github.com/aligator/haiku-golang-tools/gopls@v0.14.3
```