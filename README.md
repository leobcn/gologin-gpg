# gologin-gpg
GPG login writen in Go lang, this is a proof of concept and as such may contain major security risks

##Installation
    go get github.com/zowhoey/gologin-gpg

**this shouldn't be used on any production server as it may contain security risks**
##Flags
**-port** - port number (default "9090")

**-public_key** - GPG public keyring (default "")

**-user** - change username (default "username")

**-template** - change template location (default "login.tmpl")
