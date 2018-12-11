# Threaded.chat #
4 chan slack

## Dev Installation ##

* Install Go. Instructions can be found [here](https://golang.org/doc/install).
    * Make sure you have your `GOPATH` setup correctly. Official instructions can be found [here](https://github.com/golang/go/wiki/SettingGOPATH).
    * You might also want to add the line `export PATH=$PATH:$(go env GOPATH)/bin` to your `.bash_profile` or `.bashrc`. This is to setup Go's installation folder as a executable binary path.
* Clone this repo under `$GOPATH/src/github.com/bli940505/`
    * This is to follow official styling convention.
* Get `dep`, Go's default package mangager (still in beta) by running `brew install dep`
* Inside this folder, run `dep ensure`.
    * Nothing printed means it's successful.
