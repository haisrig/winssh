# Winssh
Ssh server implementation for windows and linux machines using Go/GoLang with password authentication.
By default SSH server run on port **2200**. Default username is **foo** and password is **bar**.
I have tried this as part of my GoLang learning and practice.


# How to Use
Get the dependent module
`go get github.com/gliderlabs/ssh`

Run
`go run main.go`

# Connecting to SSH server
While connecting to the SSH server you have to provide a command.
> ssh foo@192.168.0.106 -p 2200 dir
