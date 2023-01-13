# Getting Started 

## Install (Mac OS)

1. install via brew 
```
brew install go
```
2. Set GOPATH in bash profile 

(Since Go 1.8, the GOPATH environment variable has the default value of $HOME/go)

https://www.programming-books.io/essential/go/mac-os-install-and-setup-a223f756151c45cd9cc6acd2db0ebda0

Add the following to your ~/.bash_profile file:

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

Load the changes with source ~/.bash_profile or launch a new terminal for the changes to take effect.